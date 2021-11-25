// �ڵ��ɫ����
Vue.component("node-role-name", {
	props: ["v-role"],
	data: function () {
		let roleName = ""
		switch (this.vRole) {
			case "node":
				roleName = "��Ե�ڵ�"
				break
			case "monitor":
				roleName = "��ؽڵ�"
				break
			case "api":
				roleName = "API�ڵ�"
				break
			case "user":
				roleName = "�û�ƽ̨"
				break
			case "admin":
				roleName = "����ƽ̨"
				break
			case "database":
				roleName = "���ݿ�ڵ�"
				break
			case "dns":
				roleName = "DNS�ڵ�"
				break
			case "report":
				roleName = "�������ն�"
				break
		}
		return {
			roleName: roleName
		}
	},
	template: `<span>{{roleName}}</span>`
})