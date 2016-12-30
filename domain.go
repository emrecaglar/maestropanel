package maestropanel

type DomainExecutionResult struct {
    Result
    Details DomainOperationResult
}

type GetDomainAliasResult struct {
    Result
    Details []string
}

type DomainOperationResult struct {
    Code int32
    Message string
    Id int32
    Name string
    Username string
    Password string
    DomainUser bool
    IpString string
    ModuleResults []DomainOperationModuleResult
}

type DomainOperationModuleResult struct {
    Status bool
    Msg string
    Name string
    PType string
}

type DomainListItemResult struct {
    Result
    Details DomainListItem
}


type Domain struct {
    Name string `json:"name"`
    PlanAlias string `json:"planAlias"`
    UserName string `json:"username"`
    Password string `json:"password"`
    ActiveDomainUser bool `json:"activedomainuser"`
    FirstName string `json:"firstname"`
    LastName string `json:"lastname"`
    Email string `json:"email"`
    Expiration string `json:"expiration"`
    IPAddr string `json:"ipaddr"`
}

type SubDomain struct {
    DomainName string `json:"name"`
    SubDomain string `json:"subdomain"`
    FTPUser string `json:"ftpuser"`
}

type DomainPassword struct {
    Name string `json:"name"`
    NewPassword string `json:"newpassword"`
}

type DomainAlias struct {
    Name string `json:"name"`
    Alias string `json:"alias"`
}

type SetSubDomainFTPAccount struct {
    Name string `json:"name"`
    SubDomain string `json:"subdomain"`
    NewFTPUser string `json:"newftpuser"`
}

type GetListResult struct {
	Result
	Details []DomainListItem 
}

type GetSubDomainsResult struct {
    Result
    Details []SubDomainConfig
}

type SubDomainConfig struct {
    Name string
    FtpUser string
}

type DomainListItem struct {
    Id int32
    Name string
    ExpirationDate string
    Status int32
    OwnerName string
    Email int32
    Disk int32
}

type ChangeIpAddres struct {
    DomainName string `json:"name"`
    NewIPAddres string `json:"newip"`
}