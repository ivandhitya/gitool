package tag

const (
	KeyDataID      = "id"       //yes 	The project ID
	KeyDataOrderBy = "order_by" //no 	Return tags ordered by name or updated fields. Default is updated
	KeyDataSort    = "sort"     //no 	Return tags sorted in asc or desc order. Default is desc
	KeyDataSearch  = "search"   //no 	Return list of tags matching the search criteria. You can use ^term and term$ to find tags that begin and end with term respectively.
	KeyDataTagName = "tag_name" //yes	The name of a tag
	KeyDataRef     = "ref"      //yes	Create tag using commit SHA, another tag name, or branch name
	KeyDataMessage = "message"  //no	Creates annotated tag
)

type Sort string

const (
	SortDesc Sort = "desc"
	SortAsc  Sort = "asc"
)

type OrderBy string

const (
	OrderByName    OrderBy = "name"
	OrderByUpdated OrderBy = "updated"
)
