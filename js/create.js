window.onload = function () {
  var app = new Vue({
    el: '#main',
    data: {
      id: "",
      password: "",
    },
    created: function(){
    },
    watch: {
    },
    methods: {
        post : function() {
            var user = {};
            user.userid = this.id;
            user.pw = this.password;

            fetch("/create", {
                method: 'POST',
                mode: 'cors',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(user)
            }).then(response => {
                if (response.ok) {
                    location.href='/client/' + user.userid;
                } else if (response.status == 400) {
                    // TODO: erorr
                    console.log(response.status);
                } else {
                    console.log("NG");
                }
            });
        }
    }
  })
}
