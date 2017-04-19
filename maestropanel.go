package maestropanel

type MaestroPanel struct {
	Url     string
	Key     string
	Version string
}

func (mp *MaestroPanel) Web() Web {
	return Web{*mp}
}

func (mp *MaestroPanel) FileManager() FileManager {
	return FileManager{*mp}
}

func (mp *MaestroPanel) WebStats() WebStats {
	return WebStats{*mp}
}

func (mp *MaestroPanel) Mail() Mail {
	return Mail{*mp}
}

func (mp *MaestroPanel) Database() Database {
	return Database{*mp}
}

func (mp *MaestroPanel) FTP() FTP {
	return FTP{*mp}
}

func (mp *MaestroPanel) SSL() SSL {
	return SSL{*mp}
}

func (mp *MaestroPanel) Dns() DNS {
	return DNS{*mp}
}

func (mp *MaestroPanel) User() User {
	return User{*mp}
}
