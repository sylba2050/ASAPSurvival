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

            var shaObj = new jsSHA("SHA-256", "TEXT");
            shaObj.update(this.id);
            shaObj.update(this.password);
            shaObj.update(this.auth_code);

            user.code = shaObj.getHash("HEX");

            fetch("/login", {
                method: 'POST',
                mode: 'cors',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(user)
            }).then(response => {
                if (response.ok) {
                    document.cookie = "userid=" + user.userid;
                    document.cookie = "code=" + user.code;

                    location.href='/client/' + user.userid
                } else if (response.status == 401) {
                    // TODO: erorr表示
                    console.log(response.status);
                    location.reload();
                } else {
                    console.log("NG");
                }
            });
        },
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
