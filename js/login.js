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

            fetch("/login", {
                method: 'POST',
                mode: 'cors',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(user)
            }).then(response => {
                if (response.ok) {
                    location.href='/create'
                    console.log("success");
                } else {
                    console.log("failed");
                }
            });
        }
    }
  })
}
