package pagination

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
	"sync"

	"github.com/billowdev/go-fiber-e-commerce/pkg/configs"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// NewPagSuccessResponse constructs a paginated API success response with a message and data.
//
// This function creates a standardized success response for paginated APIs. It includes a message,
// the provided data, and pagination information. The pagination details are supplied through the
// `data` parameter, and the `newData` parameter contains the actual data to be included in the response.
//
// Parameters:
//   - message string: A custom message to include in the response status message.
//   - data PaginationInfo: An instance of PaginationInfo containing pagination details such as links, total items,
//     current page, page size, and total pages.
//   - newData interface{}: The data to be included in the response, typically a list of items or records.
//
// Returns:
// - APIV2PaginationResponse: A paginated API response with the provided message, data, and pagination information.
//
// Example:
//
//	paginationInfo := PaginationInfo{ ... } // Setup pagination info
//	items := []YourDataType{ ... }           // Data to include in response
//	response := NewPagSuccessResponse("Data retrieved successfully", paginationInfo, items)
//	// Use response in your API handler
func NewPagSuccessResponse(message string, data PaginationInfo, newData interface{}) APIV2PaginationResponse {
	return APIV2PaginationResponse{
		StatusCode:    configs.API_SUCCESS_CODE,
		StatusMessage: message,
		Data:          newData,
		Pagination: PaginationInfo{
			Links:      data.Links,
			Total:      data.Total,
			Page:       data.Page,
			PageSize:   data.PageSize,
			TotalPages: data.TotalPages,
		},
	}
}

// NewPagErrorResponse constructs a paginated API error response with a message and error details.
//
// This function creates a standardized error response for paginated APIs. It formats the error message
// by combining the provided message with the details of the error. The response includes an empty data
// set and pagination information to align with the structure of successful paginated responses.
//
// Parameters:
// - message string: A custom message to include in the response status message.
// - err error: The error that occurred, which will be included in the response message.
//
// Returns:
//   - APIV2PaginationResponse: A paginated API response with the formatted error message, an empty data
//     set, and an empty PaginationInfo struct.
//
// Example:
//
//	err := errors.New("something went wrong")
//	response := NewPagErrorResponse("Failed to retrieve data", err)
//	// Use response in your API handler
func NewPagErrorResponse(message string, err error) APIV2PaginationResponse {
	message = fmt.Sprintf("%v ERROR: %v", message, err)
	return APIV2PaginationResponse{
		StatusCode:    configs.API_SUCCESS_CODE,
		StatusMessage: message,
		Data:          []map[string]interface{}{},
		Pagination:    PaginationInfo{}, // Return an empty PaginationInfo struct
	}
}

// NewPaginationResponse constructs and sends a paginated API response in JSON format.
//
// This function creates a paginated response based on the provided data and message. It handles
// the case where there are no rows of data by returning an empty result set with appropriate
// pagination information. If there are rows of data, it returns the data along with pagination
// details in the response.
//
// Parameters:
// - c *fiber.Ctx: The Fiber context used to construct and send the HTTP response.
// - message string: A message to include in the response status message.
// - data Pagination[[]T]: A Pagination object containing the paginated data and pagination info.
//
// Returns:
// - error: An error if there was an issue sending the response, otherwise nil.
//
// Example:
//
//	data := Pagination[[]MyType]{
//	    Rows:       myData,
//	    Total:      100,
//	    Page:       1,
//	    PageSize:   10,
//	    TotalPages: 10,
//	}
//	err := NewPaginationResponse(c, "Data retrieved successfully", data)
//	if err != nil {
//	    // Handle error
//	}
func NewPaginationResponse[T any](c *fiber.Ctx, message string, data Pagination[[]T]) error {
	if data.Rows == nil || reflect.ValueOf(data.Rows).Len() == 0 {
		return c.Status(200).JSON(APIV2PaginationResponse{
			StatusCode:    configs.API_SUCCESS_CODE,
			StatusMessage: "The process of pagination was success",
			Data:          make([]interface{}, 0),
			Pagination: GetPaginationInfo(Pagination[interface{}]{
				Total:      0,
				Links:      PaginationLinks{},
				Page:       0,
				PageSize:   0,
				TotalPages: 0,
				Rows:       make([]interface{}, 0),
			}),
		})
	}
	response := APIV2PaginationResponse{
		StatusCode:    configs.API_SUCCESS_CODE,
		StatusMessage: message,
		Data:          data.Rows,
		Pagination:    GetPaginationInfo(data),
	}
	return c.Status(200).JSON(response)
}

// GetPaginationInfo extracts pagination details from a Pagination[T] payload.
//
// This function takes a Pagination[T] payload and returns a PaginationInfo object that
// includes pagination metadata such as links, total items, current page, page size,
// and total pages. It is used to simplify the extraction of pagination details from a
// paginated response.
//
// Parameters:
// - payload Pagination[T]: A Pagination[T] object containing pagination details.
//
// Returns:
//   - PaginationInfo: An object containing the pagination metadata extracted from the
//     provided payload.
//
// Example:
//
//	paginatedPayload := Pagination[MyType]{
//	    Links: PaginationLinks{ Next: "next-url", Previous: "prev-url" },
//	    Total: 100,
//	    Page:  2,
//	    PageSize: 10,
//	    TotalPages: 10,
//	}
//	info := GetPaginationInfo(paginatedPayload)
//	// info will be a PaginationInfo object with the same values as paginatedPayload
func GetPaginationInfo[T any](payload Pagination[T]) PaginationInfo {
	return PaginationInfo{
		Links:      payload.Links,
		Total:      payload.Total,
		Page:       payload.Page,
		PageSize:   payload.PageSize,
		TotalPages: payload.TotalPages,
	}
}

// GetOffset calculates the offset for pagination based on the current page and limit.
//
// This method computes the offset value used in database queries to fetch the correct subset
// of records for the current page. The offset is calculated as the number of records to skip
// before the current page, based on the page number and the number of records per page.
//
// Returns:
//   - int - The offset value used for pagination queries. This is calculated as
//     (current page number - 1) * records per page.
//
// Example:
//
//	params := PaginationParams[MyFilterType]{ Page: 2, Limit: 10 }
//	offset := params.GetOffset() // offset will be 10
func (p *PaginationParams[FilterType]) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

// GetPage returns the current page number for pagination.
//
// This method provides the page number for pagination, ensuring that it defaults to 1 if
// the page number is not set (i.e., it is zero). This is useful for handling pagination
// logic where a default page number is required if none is specified.
//
// Returns:
// - int - The current page number. Defaults to 1 if the Page field is 0.
//
// Example:
//
//	params := PaginationParams[MyFilterType]{ Page: 0 }
//	page := params.GetPage() // page will be 1
func (p *PaginationParams[FilterType]) GetLimit() int {
	if p.Limit == 0 {
		return 10
	}
	return p.Limit
}

// GetPage returns the current page with a default value if not set.
func (p *PaginationParams[FilterType]) GetPage() int {
	if p.Page == 0 {
		return 1
	}
	return p.Page
}

// GetSort returns the sorting order for the pagination parameters.
//
// This method provides the sorting criteria used when querying data. If no specific sorting
// order has been set in the `PaginationParams`, it defaults to sorting by "Id" in descending
// order.
//
// Returns:
//   - string - The sorting order for the query. If `p.Sort` is empty, returns "Id desc" as the default
//     sorting order; otherwise, returns the value of `p.Sort`.
//
// Note:
//   - The default sorting order is "Id desc". If a different sorting order is required, ensure
//     that `p.Sort` is set appropriately before calling this method.
func (p *PaginationParams[FilterType]) GetSort() string {
	if p.Sort == "" {
		return "id DESC"
	}
	return p.Sort
}

// GetAPIEndpoint constructs the full API endpoint URL based on the current request context.
//
// This function retrieves the original URL of the request, omitting any query string, and
// constructs a new URL by combining the current host with the path of the request. It is
// intended to generate a full URL to represent the endpoint being accessed.
//
// Parameters:
// - c: *fiber.Ctx - The Fiber context that provides access to the request details.
//
// Returns:
// - string - The full API endpoint URL, including the scheme (https), host, and path.
//
// The function performs the following steps:
//  1. Retrieves the original URL of the request using `c.OriginalURL()`, which includes both
//     the path and the query string.
//  2. Removes the query string from the URL if it exists by locating the "?" character and
//     slicing the URL up to that index.
//  3. Constructs the full URL using the "https" scheme, the host from `c.Hostname()`, and
//     the path obtained in the previous step.
//
// Note: The function assumes that the scheme should always be "https". If the scheme might
// vary or should be configurable, additional logic should be added.
func GetAPIEndpoint(c *fiber.Ctx) string {
	// Get the original URL from the request (excluding the query string)
	originalURL := c.OriginalURL()

	// Remove the query string (if any)
	if index := strings.Index(originalURL, "?"); index != -1 {
		originalURL = originalURL[:index]
	}

	// Construct the desired URL with the current host and path
	return fmt.Sprintf("https://%s%s", c.Hostname(), originalURL)
}

// Paginate retrieves a paginated set of records based on the provided pagination parameters.
//
// This function handles pagination by calculating the total number of rows and pages,
// and then retrieves a specific page of records from the database. It generates navigation
// links for paginated data, such as links to the next and previous pages, based on the
// current page and the total number of pages.
//
// Parameters:
// - p: PaginationParams[FT] - A struct containing pagination parameters such as the current page number, page size, sorting options, etc.
// - query: *gorm.DB - The GORM database query object that defines the dataset being paginated.
//
// Returns:
// - Pagination[T] - A struct that contains the paginated data, total rows, page information, and pagination links.
// - error - An error object if an error occurs during the query.
//
// The function first counts the total number of records to determine the total number of pages.
// It then fetches the records for the current page using the offset and limit specified in
// the pagination parameters. It also constructs next and previous page links based on the
// current page and total pages for ease of navigation.
//
// Note: This function retrieves a single page of records and assumes that the `Rows` field
// in the `Pagination` struct can hold a single record of type `T`. This may need to be adjusted
// depending on the actual data structure and requirements.
func Paginate[FT any, T any](p PaginationParams[FT], query *gorm.DB) (Pagination[T], error) {
	var value T
	var totalRows int64
	if err := query.Model(value).Count(&totalRows).Error; err != nil {
		return Pagination[T]{}, err
	}
	p.TotalRows = totalRows
	p.TotalPages = int(math.Ceil(float64(totalRows) / float64(p.GetLimit())))

	if err := query.Offset(p.GetOffset()).Limit(p.GetLimit()).Order(p.GetSort()).Find(&value).Error; err != nil {
		return Pagination[T]{}, err

	}
	var nextLink, prevLink string
	if p.Page < p.TotalPages {
		nextLink = fmt.Sprintf("%s?page=%d&page_size=%d", p.BaseURL, p.Page+1, p.Limit)
	}
	if p.Page > 1 {
		prevLink = fmt.Sprintf("%s?page=%d&page_size=%d", p.BaseURL, p.Page-1, p.Limit)
	}
	return Pagination[T]{
		Links: PaginationLinks{
			Next:     nextLink,
			Previous: prevLink,
		},
		Total:      p.TotalRows,
		Page:       p.GetPage(),
		PageSize:   p.GetLimit(),
		TotalPages: p.TotalPages,
		Rows:       value,
	}, nil
}

// PaginateBatchProcessing performs pagination by processing records in batches based on the provided query.
//
// This function handles paginated fetching of records from a database using GORM. If the number of records to
// be fetched (defined by the pagination limit) exceeds the specified query limiter, the function will process
// records in batches. Otherwise, it fetches all records in a single query. It also generates links for
// navigating to the next and previous pages.
//
// Parameters:
// - p: PaginationParams[F] - Contains the pagination parameters such as the current page, page size, and sort order.
// - query: *gorm.DB - The GORM database query used to fetch the records.
// - queryLimitor: int - A threshold for determining whether batch processing should be applied.
// - batchSize: int - The number of records to be fetched per batch when batch processing is applied.
//
// Returns:
// - PaginationBatchProcessingResponse[T] - A struct containing the paginated result set, pagination links, and other details.
// - error - If any error occurs during querying or batch processing, it is returned.
//
// The function accumulates the results from each batch if batch processing is applied, and returns the full set
// of records along with pagination details like total rows, total pages, and links to navigate between pages.
func PaginateBatchProcessing[F any, T any](p PaginationParams[F], query *gorm.DB, queryLimitor int, batchSize int) (PaginationBatchProcessingResponse[T], error) {
	var totalRows int64
	var results []T // Declare results as a slice of T

	// Count total rows based on the query
	if err := query.Model(new(T)).Count(&totalRows).Error; err != nil {
		return PaginationBatchProcessingResponse[T]{}, err
	}
	p.TotalRows = totalRows
	p.TotalPages = int(math.Ceil(float64(totalRows) / float64(p.GetLimit())))

	// If limit is more than limitor, apply batch processing
	if p.GetLimit() > queryLimitor {
		totalBatches := int(math.Ceil(float64(p.GetLimit()) / float64(batchSize)))

		// Process data in batches
		for batch := 0; batch < totalBatches; batch++ {
			var batchResults []T // Declare batchResults as a slice of T

			// Calculate the offset for each batch
			offset := batch * batchSize
			limit := batchSize

			// If it's the last batch and we don't need a full batch size
			if batch == totalBatches-1 {
				limit = p.GetLimit() - (batch * batchSize)
			}

			// Fetch the batch of records
			if err := query.Offset(offset).Limit(limit).Order(p.GetSort()).Find(&batchResults).Error; err != nil {
				return PaginationBatchProcessingResponse[T]{}, err
			}

			// Append batch results to the main result set
			results = append(results, batchResults...)
		}
	} else {
		// If limit <= limitor, handle normally with a single query
		if err := query.Offset(p.GetOffset()).Limit(p.GetLimit()).Order(p.GetSort()).Find(&results).Error; err != nil {
			return PaginationBatchProcessingResponse[T]{}, err
		}
	}

	// Generate next and previous links for pagination
	var nextLink, prevLink string
	if p.Page < p.TotalPages {
		nextLink = fmt.Sprintf("%s?page=%d&page_size=%d", p.BaseURL, p.Page+1, p.Limit)
	}
	if p.Page > 1 {
		prevLink = fmt.Sprintf("%s?page=%d&page_size=%d", p.BaseURL, p.Page-1, p.Limit)
	}

	// Return the complete paginated result set
	return PaginationBatchProcessingResponse[T]{
		Links: PaginationLinks{
			Next:     nextLink,
			Previous: prevLink,
		},
		Total:      p.TotalRows,
		Page:       p.GetPage(),
		PageSize:   p.GetLimit(),
		TotalPages: p.TotalPages,
		Rows:       results, // Return the accumulated results from all batches
	}, nil
}

// PaginateBatchProcessingSync processes paginated data in batches using goroutines for concurrent fetching and processing.
//
// This function fetches and processes records from a database query in batches, allowing for efficient handling of large datasets.
// It uses goroutines to parallelize the fetching of batches to improve performance. The function also calculates and provides
// pagination details such as total rows, total pages, and links for the next and previous pages.
//
// Parameters:
// - p: PaginationParams[F] - An instance of PaginationParams that contains pagination settings such as current page, page size, and sorting order.
// - query: *gorm.DB - A GORM database query object used to fetch data from the database.
// - queryLimitor: int - A limit for the number of concurrent queries to prevent overwhelming the database or application resources.
// - batchSize: int - The number of records to fetch in each batch.
//
// Returns:
// - PaginationBatchProcessingResponse[T] - A response containing the paginated data, pagination details, and any potential errors.
// - error - An error if any occurs during the batch processing or querying.
//
// The function handles errors that occur during the fetching of records and ensures that all data is collected and returned in a
// structured format, including pagination links for navigation between pages.
func PaginateBatchProcessingSync[F any, T any](p PaginationParams[F], query *gorm.DB, queryLimitor int, batchSize int) (PaginationBatchProcessingResponse[T], error) {
	var totalRows int64
	var results []T // Declare results as a slice of T

	// Count total rows based on the query
	if err := query.Model(new(T)).Count(&totalRows).Error; err != nil {
		return PaginationBatchProcessingResponse[T]{}, err
	}
	p.TotalRows = totalRows
	p.TotalPages = int(math.Ceil(float64(totalRows) / float64(p.GetLimit())))

	// Channel for results and error handling
	resultsChan := make(chan []T, p.TotalPages) // Channel to collect results from goroutines
	errChan := make(chan error, 1)              // Channel to report errors

	var wg sync.WaitGroup // WaitGroup to wait for all goroutines to complete

	// Calculate totalBatches here
	totalBatches := int(math.Ceil(float64(p.GetLimit()) / float64(batchSize)))

	// Function to process a single batch
	processBatch := func(batch int) {
		defer wg.Done() // Signal the wait group that this batch is done

		var batchResults []T // Declare batchResults as a slice of T
		offset := batch * batchSize
		limit := batchSize
		if batch == totalBatches-1 {
			limit = p.GetLimit() - (batch * batchSize)
		}

		// Fetch the batch of records
		if err := query.Offset(offset).Limit(limit).Order(p.GetSort()).Find(&batchResults).Error; err != nil {
			errChan <- err
			return
		}
		resultsChan <- batchResults
	}

	if p.GetLimit() > queryLimitor {
		for batch := 0; batch < totalBatches; batch++ {
			wg.Add(1)
			go processBatch(batch)
		}

		// Close resultsChan after all batches are processed
		go func() {
			wg.Wait()
			close(resultsChan)
		}()

		// Collect results from the channel
		for batchResults := range resultsChan {
			results = append(results, batchResults...)
		}

		// Check for errors
		select {
		case err := <-errChan:
			return PaginationBatchProcessingResponse[T]{}, err
		default:
			// No errors, continue
		}
	} else {
		if err := query.Offset(p.GetOffset()).Limit(p.GetLimit()).Order(p.GetSort()).Find(&results).Error; err != nil {
			return PaginationBatchProcessingResponse[T]{}, err
		}
	}

	// Generate next and previous links for pagination
	var nextLink, prevLink string
	if p.Page < p.TotalPages {
		nextLink = fmt.Sprintf("%s?page=%d&page_size=%d", p.BaseURL, p.Page+1, p.Limit)
	}
	if p.Page > 1 {
		prevLink = fmt.Sprintf("%s?page=%d&page_size=%d", p.BaseURL, p.Page-1, p.Limit)
	}

	return PaginationBatchProcessingResponse[T]{
		Links: PaginationLinks{
			Next:     nextLink,
			Previous: prevLink,
		},
		Total:      p.TotalRows,
		Page:       p.GetPage(),
		PageSize:   p.GetLimit(),
		TotalPages: p.TotalPages,
		Rows:       results,
	}, nil
}

// NewPaginationParams creates a new instance of PaginationParams from the given Fiber context.
//
// This function extracts pagination parameters (limit, page, sort, and order) from the query
// parameters of the request context. It initializes the pagination parameters with default values
// and updates them based on the provided query parameters if they are valid.
//
// Parameters:
// - c: *fiber.Ctx - The Fiber context containing the HTTP request with query parameters.
//
// Returns:
//   - PaginationParams[FilterType] - An instance of PaginationParams populated with values from
//     the query parameters or default values if the query parameters are not provided or invalid.
//
// Notes:
//   - Default values are used if the query parameters are missing or invalid: Limit is set to 10,
//     Page is set to 1, and Sort is set to "created_at desc".
//   - If the "limit" query parameter is provided and is a positive integer, it overrides the default limit.
//   - If the "page" query parameter is provided and is a positive integer, it overrides the default page.
//   - If the "sort" query parameter is provided and is non-empty, it overrides the default sort value.
//   - If the "order" query parameter is provided and is non-empty, it sets the Order field of the pagination parameters.
//
// Example:
//
//	c := fiber.New().Ctx()
//	params := NewPaginationParams[MyFilterType](c)
//	// params now contains pagination parameters extracted from the request query or default values.
func NewPaginationParams[FilterType interface{}](c *fiber.Ctx) PaginationParams[FilterType] {
	host := GetAPIEndpoint(c) // Assuming utils.GetAPIEndpoint is a function you have
	defaultLimit := 10
	defaultPage := 1
	defaultSort := "created_at desc"

	paginationParams := PaginationParams[FilterType]{
		Limit:   defaultLimit,
		Page:    defaultPage,
		Sort:    defaultSort,
		BaseURL: host,
	}

	limit, err := strconv.Atoi(c.Query("limit"))
	if err == nil && limit > 0 {
		paginationParams.Limit = limit
	}

	page, err := strconv.Atoi(c.Query("page"))
	if err == nil && page > 0 {
		paginationParams.Page = page
	}

	sort := c.Query("sort")
	if sort != "" {
		paginationParams.Sort = sort
	}

	order := c.Query("order")
	if sort != "" {
		paginationParams.Order = order
	}
	return paginationParams
}

// PaginateArray performs pagination on a slice of data and provides the corresponding
// pagination information and the subset of data for the requested page.
//
// This function slices the provided data according to the specified page number and page size,
// and returns the relevant data along with pagination information. It also constructs next
// and previous links for navigating between pages based on the provided endpoint URL.
//
// Parameters:
// - data: []T - The slice of data to be paginated.
// - page: int - The current page number, starting from 1.
// - pageSize: int - The number of items per page.
// - endpoint: string - The base URL for constructing pagination links.
//
// Returns:
//   - PaginationInfo - Struct containing pagination metadata such as links, total items,
//     current page, page size, and total pages.
//   - []T - The subset of data for the specified page.
//
// Errors:
//   - If `pageSize` or `page` is less than or equal to 0, or if `page` exceeds the total number
//     of pages, the function returns an empty `PaginationInfo` and `nil` data. Consider returning
//     an error for more robust error handling in production code.
//
// Notes:
//   - The function calculates the total number of pages based on the total number of items and
//     `pageSize`. It adjusts the slice boundaries to ensure that the end index does not exceed
//     the length of the data slice.
//   - Pagination links are constructed using the `endpoint` URL and the current page and page size.
//     If `page` is beyond the available pages, it handles it by returning empty results.
func PaginateArray[T any](data []T, page, pageSize int, endpoint string) (PaginationInfo, []T) {
	totalItems := len(data)
	if pageSize <= 0 || page <= 0 {
		return PaginationInfo{}, nil // or return an error
	}

	totalPages := (totalItems + pageSize - 1) / pageSize

	if page > totalPages {
		return PaginationInfo{}, nil // or return an error
	}

	start := (page - 1) * pageSize
	end := start + pageSize
	if end > totalItems {
		end = totalItems
	}

	pageData := data[start:end]

	// Construct next and previous links
	var nextLink, prevLink string
	if page < totalPages {
		nextLink = fmt.Sprintf("%s?page=%d&page_size=%d", endpoint, page+1, pageSize)
	}
	if page > 1 {
		prevLink = fmt.Sprintf("%s?page=%d&page_size=%d", endpoint, page-1, pageSize)
	}

	pagination := PaginationInfo{
		Links: PaginationLinks{
			Next:     nextLink,
			Previous: prevLink,
		},
		Total:      int64(totalItems),
		Page:       page,
		PageSize:   pageSize,
		TotalPages: totalPages,
	}
	return pagination, pageData
}
