window.onload = function () {
  var app = new Vue({
    el: '#main',
    data: {
      is_join: {},
      is_survival: {},
      joins: {},
      survivals: {},
      chats: {},
      joins_userid: [],
      survivals_userid: [],
      count: 0,
      userid: "",
      team: "yellow",
      auth_code: "",
      active_chat: "default",
      is_dead_or_resporn_active: false,
    },
    created: function(){
        setInterval(() => { this.count++ }, 1000)

        var cookie = this.getCookie();
        this.userid = cookie["userid"];
        this.auth_code = cookie["code"];

        this.getChat();

        this.isSurvival();
        this.isJoin();
        this.getSurvivals();
        this.getJoins();
    },
    watch: {
        count: function(){
            this.getSurvivals();
            this.getJoins();
            this.getChat();
        },
        joins: function(){
            var data = [];
            for (var i in this.joins) {
                // TODO DB側で処理
                if(this.joins[i].userid == this.userid) continue;
                data.push(this.joins[i].userid);
            }
            this.joins_userid = data;
        },
        survivals: function(){
            var data = [];
            for(var i in this.survivals) {
                // TODO DB側で処理
                if(this.survivals[i].userid == this.userid) continue;
                data.push(this.survivals[i].userid);
            }
            this.survivals_userid = data;
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
        dead : function() {
            fetch("/dead/" + this.userid + "?code=" + this.auth_code, {
                method: 'POST',
                mode: 'cors',
            }).then(response => {
                if (response.ok) {
                    location.reload();
                    console.log("OK");
                } else if (response.status == 307) {
                    location.href='/'
                } else {
                    // TODO: erorr表示
                    console.log("NG");
                }
            });
        },
        resporn : function() {
            fetch("/resporn/" + this.userid + "?code=" + this.auth_code, {
                method: 'POST',
                mode: 'cors',
            }).then(response => {
                if (response.ok) {
                    location.reload();
                    console.log("OK");
                } else if (response.status == 307) {
                    location.href='/'
                } else {
                    // TODO: erorr表示
                    console.log("NG");
                }
            });
        },
        join : function() {
            fetch("/join/" + this.userid + "?code=" + this.auth_code, {
                method: 'POST',
                mode: 'cors',
            }).then(response => {
                if (response.ok) {
                    location.reload();
                    console.log("OK");
                } else if (response.status == 307) {
                    location.href='/'
                } else {
                    // TODO: erorr表示
                    console.log("NG");
                }
            });
        },
        breaktime : function() {
            fetch("/dontjoin/" + this.userid + "?code=" + this.auth_code, {
                method: 'POST',
                mode: 'cors',
            }).then(response => {
                if (response.ok) {
                    console.log("OK");
                    location.reload();
                } else if (response.status == 307) {
                    location.href='/'
                } else {
                    // TODO: erorr表示
                    console.log("NG");
                }
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
        checkSurvivalById : function(id) {
            return (this.survivals_userid.indexOf(id) >= 0)
        },
        setActiveChat : function(chatCode) {
            this.active_chat = chatCode
        },
        switchDeadOrRespornButtonActive : function(chatCode) {
            this.is_dead_or_resporn_active = !this.is_dead_or_resporn_active
        },
        deadOrRespornFromNowStatus : function() {
            if (this.is_survival.is_survival) {
                this.dead();
            } else {
                this.resporn();
            }
        },
        joinOrBreakTimeFromNowStatus : function() {
            if (this.is_join.is_join) {
                this.breaktime();
            } else {
                this.join();
            }
        },
        sendChat : function(content) {
            var chat = {};
            chat.content = content;
            chat.target = this.userid;
            chat.target_team = this.team;
            chat.send_user_id = this.userid;

            fetch("/chat?userid=" + this.userid + "&code=" + this.auth_code, {
                method: 'POST',
                mode: 'cors',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(chat)
            }).then(response => {
                if (response.ok) {
                    console.log("OK");
                } else if (response.status == 307) {
                    location.href='/'
                } else {
                    // TODO: erorr表示
                    console.log("NG");
                }
            });
        },
        getChat : function() {
            fetch("/chat?userid=" + this.userid + "&code=" + this.auth_code, {
                method: 'GET',
                mode: 'cors',
            }).then((response) => {
                return response.json();
            }).then((json) => {
                this.chats = json;
                this.chats = this.chats.reverse();
            });
        },
        redirect_by_url : function(url) {
            location.href = url;
        },
    },
  })
}
