$.validator.addMethod('ipv4', function(value) {
    if (value.length == 0) return true;

    var ipv4      = /^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$/;

    if (value.match(ipv4)) {
        return true;
    }
}, 'msg_invalid_ipv4');


$.validator.addMethod('ipv4_cidr', function(value) {
    if (value.length == 0) return true;

    var ipv4      = /^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$/;
    var ipv4_cidr = /^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])(\/([0-9]|[1-2][0-9]|3[0-2]))$/;

    if (value.match(ipv4) || value.match(ipv4_cidr)) {
        return true;
    }
}, 'msg_invalid_ipv4_cidr');


$.validator.addMethod('password_rule', function(value) {
    if (value.length == 0) return true;
    var regex = /^(?=.*[a-zA-Z])(?=.*[!@#$%^*+=-])(?=.*[0-9]).{9,16}$/;
    //var regex = /^(?=.*[a-z])(?=.*[A-Z])(?=.*[!@#$%^*+=-])(?=.*[0-9]).{9,16}$/;
    if (value.match(regex)) {
        return true;
    }
}, 'about_password');


$.validator.addMethod("notEqual", function(value, element, param) {
	return this.optional(element) || value != $(param).val();
}, "This has to be different...");

