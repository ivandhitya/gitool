package commit

const (
	KeyDataID          = "id"           //yes 	The project ID
	KeyDataRefName     = "ref_name"     //string	no	The name of a repository branch, tag or revision range, or if not given the default branch
	KeyDataSince       = "since"        //string	no	Only commits after or on this date are returned in ISO 8601 format YYYY-MM-DDTHH:MM:SSZ
	KeyDataUntil       = "until"        //string	no	Only commits before or on this date are returned in ISO 8601 format YYYY-MM-DDTHH:MM:SSZ
	KeyDataPath        = "path"         //string	no	The file path
	KeyDataAll         = "all"          //boolean	no	Retrieve every commit from the repository
	KeyDataWithStats   = "with_stats"   //boolean	no	Stats about each commit are added to the response
	KeyDataFirstParent = "first_parent" //boolean	no	Follow only the first parent commit upon seeing a merge commit
	KeyDataOrder       = "order"        //string	no	List commits in order. Possible values: default, topo. Defaults to default, the commits are shown in reverse chronological order.
	KeyDataTrailers    = "trailers"     //boolean	no	Parse and include Git trailers for every commit
)
