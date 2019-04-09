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

        var cookie = this.getCookie();
        this.userid = cookie["userid"];
        this.auth_code = cookie["code"];
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
        getCookie : function() {
            var arr = new Array();
            if(document.cookie != ''){
                var tmp = document.cookie.split('; ');
                for(var i=0;i<tmp.length;i++){
                    var data = tmp[i].split('=');
                    arr[data[0]] = decodeURIComponent(data[1]);
                }
            }
            return arr;
        },
    },
  })
}
