package ikuai

import (
	"github.com/jakeslee/ikuai/action"
	"testing"
)

func getClient() *IKuai {
	i := NewIKuai("http://ikuai.com", "admin", "hello12345", true, true)

	return i
}

func TestLogin(t *testing.T) {

	login, err := getClient().Login()
	if err != nil {
		t.Error(err)
	}

	t.Log(login)
}

func TestIKuai_ShowIPGroup(t *testing.T) {
	client := getClient()

	login, err := client.Login()
	if err != nil {
		t.Error(err)
	}

	res := &action.ShowIPGroupResult{}

	ip, err := client.Run(login, action.NewIPGroupShowAction(), res)
	if err != nil {
		t.Error(err)
	}

	t.Log(ip)
}

func TestIKuai_ShowMonitorLan(t *testing.T) {
	client := getClient()

	_, err := client.Login()
	if err != nil {
		t.Error(err)
	}

	result, err := client.ShowMonitorLan()
	if err != nil {
		return
	}

	if err != nil {
		t.Error(err)
	}

	t.Logf("%+v", result)
}

func TestIKuai_ShowSysStat(t *testing.T) {
	client := getClient()

	_, err := client.Login()
	if err != nil {
		t.Error(err)
	}

	result, err := client.ShowSysStat()
	if err != nil {
		return
	}

	if err != nil {
		t.Error(err)
	}

	t.Logf("%+v", result)
}

func TestIKuai_ShowLanList(t *testing.T) {
	client := getClient()

	_, err := client.Login()
	if err != nil {
		t.Error(err)
	}

	result, err := client.ShowLanList()
	if err != nil {
		return
	}

	if err != nil {
		t.Error(err)
	}

	t.Logf("%+v", result)

	cresult, err := client.ShowWanVlanInterface("wan1")
	if err != nil {
		return
	}

	if err != nil {
		t.Error(err)
	}

	t.Logf("%+v", cresult)

}
