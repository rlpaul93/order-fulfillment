package in

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/rlpaul93/order-fulfillment/internal/domain/model"
	"github.com/rlpaul93/order-fulfillment/internal/domain/service"
)

// mockPackRepository is a mock implementation of port.PackRepository for testing.
type mockPackRepository struct {
	packs []*model.Pack
	err   error
}

func (m *mockPackRepository) Create(pack *model.Pack) error {
	return m.err
}

func (m *mockPackRepository) GetByID(id uuid.UUID) (*model.Pack, error) {
	return nil, m.err
}

func (m *mockPackRepository) Update(pack *model.Pack) error {
	return m.err
}

func (m *mockPackRepository) Delete(id uuid.UUID) error {
	return m.err
}

func (m *mockPackRepository) ListByProduct(productID uuid.UUID) ([]*model.Pack, error) {
	return m.packs, m.err
}

func TestPackFulfillmentHandler_Success(t *testing.T) {
	productID := uuid.New()
	mockRepo := &mockPackRepository{
		packs: []*model.Pack{
			{ID: uuid.New(), ProductID: productID, Size: 250},
			{ID: uuid.New(), ProductID: productID, Size: 500},
			{ID: uuid.New(), ProductID: productID, Size: 1000},
		},
	}
	packSvc := &service.PackService{Repo: mockRepo}
	fulfillSvc := &service.PackFulfillmentService{}

	handler := PackFulfillmentHandler(fulfillSvc, packSvc)

	req := httptest.NewRequest(http.MethodGet, "/fulfill?product_id="+productID.String()+"&quantity=251", nil)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, rec.Code)
	}

	var result service.PackFulfillmentResult
	if err := json.NewDecoder(rec.Body).Decode(&result); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if result.TotalItems < 251 {
		t.Errorf("expected TotalItems >= 251, got %d", result.TotalItems)
	}
}

func TestPackFulfillmentHandler_InvalidProductID(t *testing.T) {
	packSvc := &service.PackService{Repo: &mockPackRepository{}}
	fulfillSvc := &service.PackFulfillmentService{}

	handler := PackFulfillmentHandler(fulfillSvc, packSvc)

	req := httptest.NewRequest(http.MethodGet, "/fulfill?product_id=invalid&quantity=100", nil)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected status %d, got %d", http.StatusBadRequest, rec.Code)
	}
}

func TestPackFulfillmentHandler_InvalidQuantity(t *testing.T) {
	productID := uuid.New()
	packSvc := &service.PackService{Repo: &mockPackRepository{}}
	fulfillSvc := &service.PackFulfillmentService{}

	handler := PackFulfillmentHandler(fulfillSvc, packSvc)

	req := httptest.NewRequest(http.MethodGet, "/fulfill?product_id="+productID.String()+"&quantity=abc", nil)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected status %d, got %d", http.StatusBadRequest, rec.Code)
	}
}

func TestPackFulfillmentHandler_NoPacksFound(t *testing.T) {
	productID := uuid.New()
	mockRepo := &mockPackRepository{
		packs: []*model.Pack{},
	}
	packSvc := &service.PackService{Repo: mockRepo}
	fulfillSvc := &service.PackFulfillmentService{}

	handler := PackFulfillmentHandler(fulfillSvc, packSvc)

	req := httptest.NewRequest(http.MethodGet, "/fulfill?product_id="+productID.String()+"&quantity=100", nil)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Errorf("expected status %d, got %d", http.StatusNotFound, rec.Code)
	}
}

func TestPackFulfillmentHandler_ExactPackSize(t *testing.T) {
	productID := uuid.New()
	mockRepo := &mockPackRepository{
		packs: []*model.Pack{
			{ID: uuid.New(), ProductID: productID, Size: 250},
			{ID: uuid.New(), ProductID: productID, Size: 500},
		},
	}
	packSvc := &service.PackService{Repo: mockRepo}
	fulfillSvc := &service.PackFulfillmentService{}

	handler := PackFulfillmentHandler(fulfillSvc, packSvc)

	req := httptest.NewRequest(http.MethodGet, "/fulfill?product_id="+productID.String()+"&quantity=500", nil)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, rec.Code)
	}

	var result service.PackFulfillmentResult
	if err := json.NewDecoder(rec.Body).Decode(&result); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if result.TotalItems != 500 {
		t.Errorf("expected TotalItems = 500, got %d", result.TotalItems)
	}
}
