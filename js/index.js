window.onload = function () {
  var app = new Vue({
    el: '#main',
    data: {
      is_join: {},
      is_survival: {},
      joins: {},
      survivals: {},
      count: 0,
      userid: "",
      auth_code: "",
    },
    created: function(){
        setInterval(() => { this.count++ }, 1000)
        this.userid = document.getElementById("userid").innerHTML;
        this.auth_code = document.getElementById("auth_code").innerHTML;
        this.isSurvival();
        this.isJoin();
        this.getSurvivals();
        this.getJoins();
    },
    watch: {
        count: function(){
            console.log(this.joins);
            console.log(this.survivals);
            console.log(this.is_join);
            console.log(this.is_survival);
        },
    },
    methods: {
        isSurvival : function() {
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
        getSurvivals : function() {
            fetch("/survival" + "?userid=" + this.userid + "&code=" + this.auth_code, {
                method: 'GET',
                mode: 'cors',
            }).then((response) => {
                return response.json();
            }).then((json) => {
                console.log(json)
                this.survivals = json;
            });
        },
        isJoin : function() {
            fetch("/join/" + this.userid + "?code=" + this.auth_code, {
                method: 'GET',
                mode: 'cors',
            }).then((response) => {
                return response.json();
            }).then((json) => {
                this.is_join = json;
            });
        },
        getJoins : function() {
            fetch("/join" + "?userid=" + this.userid + "&code=" + this.auth_code, {
                method: 'GET',
                mode: 'cors',
            }).then((response) => {
                return response.json();
            }).then((json) => {
                this.joins = json;
            });
        },
    },
  })
}
