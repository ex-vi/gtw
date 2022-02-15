package gtw

type User struct {
	ID       *string `json:"id"`
	Name     *string `json:"name"`
	Username *string `json:"username"`
}

type PartialError struct {
	ResourceType *string `json:"resource_type"`
	Field        *string `json:"field"`
	Parameter    *string `json:"parameter"`
	ResourceId   *string `json:"resource_id"`
	Title        *string `json:"title"`
	Section      *string `json:"section"`
	Detail       *string `json:"detail"`
	Value        *string `json:"value"`
	Type         *string `json:"type"`
}

type PaginationMeta struct {
	ResultCount   *int    `json:"result_count"`
	NextToken     *string `json:"next_token,omitempty"`
	PreviousToken *string `json:"previous_token,omitempty"`
}
