package maestropanel

type action struct{
	URL string
	Method string
}

var getListAction = action{"/domain/getlist", "GET"}
var createDomainAction = action{"/domain/create", "POST"}
var deleteDomainAction = action{"/domain/delete", "DELETE"}
var stopDomainAction = action{"/domain/stop", "POST"}
var passwordChangeDomainAction = action{"/domain/password", "POST"}
var addDominAliasAction = action{"/domain/adddomainalias", "POST"}
var deleteDominAliasAction = action{"/domain/deletedomainalias", "DELETE"}
var getDomainAliasesAction = action{"/domain/getdomainaliases", "GET"}
var addSubdomainAction = action{"/domain/addsubdomain", "POST"}
var deleteSubdomainAction = action{"/domain/deletesubdomain", "DELETE"}
var subDomainsAction = action{"/domain/getsubdomains","GET"}
var setSubDomainFTPAccountAction = action{"/domain/setsubdomainftpaccount","POST"}
var changeIPAddressAction = action{"/domain/changeipaddr", "POST"}
var getDomainListItemAction = action{"/domain/getlistitem", "GET"}
var getLimitsAction = action{"/domain/getlimits", "GET"}
var forwardingAction = action{"/domain/forwarding", "POST"}
var changeResellerAction = action{"/domain/changereseller", "POST"}
var setDomainPlanAction = action{"/domain/setdomainplan", "POST"}
var changeNETRuntimeVersionAction = action{"/domain/changedotnetruntimeversion", "POST"}
var getNETRuntimeVersionAction = action{"/domain/getdotnetruntimeversion", "GET"}


var setWriteAccessAction = action{"/domain/setwriteaccess", "POST"}
var revokeWriteAccessAction = action{"/domain/revokewriteaccess","POST"}
var createDirectoryAction = action{"/domain/createdirectory","POST"}
var getItemsAction = action{"/domain/getitems","GET"}