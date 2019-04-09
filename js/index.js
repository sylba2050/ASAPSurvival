window.onload = function () {
  var app = new Vue({
    el: '#main',
    data: {
      is_join: {},
      is_survival: {},
      count: 0,
      userid: "",
      auth_code: "",
    },
    created: function(){
        setInterval(() => { this.count++ }, 1000)
        this.userid = document.getElementById("userid").innerHTML;
        this.auth_code = document.getElementById("auth_code").innerHTML;
        this.getSurvival();
        this.getJoin();
    },
    watch: {
        count: function(){
        },
    },
    methods: {
        getSurvival : function() {
            fetch("/survival/" + this.userid + "?code=" + this.auth_code, {
                method: 'GET',
                mode: 'cors',
            }).then((response) => {
                return response.json();
            }).then((json) => {
                console.log(json)
                this.is_survival = json;
            });
        },
        getJoin : function() {
            fetch("/join/" + this.userid + "?code=" + this.auth_code, {
                method: 'GET',
                mode: 'cors',
            }).then((response) => {
                return response.json();
            }).then((json) => {
                this.is_join = json;
            });
        },
    },
  })
}
