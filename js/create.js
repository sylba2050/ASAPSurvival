window.onload = function () {
  var app = new Vue({
    el: '#main',
    data: {
      id: "",
      password: "",
      auth_code: "",
    },
    created: function(){
    },
    watch: {
        auth_code: function() {
            var user = {};
            user.userid = this.id;
            user.pw = this.password;

            var shaObj = new jsSHA("SHA-256", "TEXT");
            shaObj.update(this.id);
            shaObj.update(this.password);
            shaObj.update(this.auth_code);

            var code = shaObj.getHash("HEX");

            fetch("/create", {
                method: 'POST',
                mode: 'cors',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(user)
            }).then(response => {
                if (response.ok) {
                    document.cookie = "userid=" + user.userid;
                    document.cookie = "code=" + code;

                    location.href='/client/' + user.userid
                } else if (response.status == 400) {
                    // TODO: erorr
                    console.log(response.status);
                    location.reload();
                } else {
                    console.log("NG");
                }
            });
        }
    },
    methods: {
        post : function() {
            fetch("/code/" + this.id, {
                method: 'GET',
                mode: 'cors',
            }).then(response => {
                if (response.ok) {
                    return response.text()
                } else {
                    console.log("NG");
                    throw new Error();
                }
            }).then(text =>{
                this.auth_code = text
            });
        }
    }
  })
}
