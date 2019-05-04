window.onload = function () {
  var app = new Vue({
    el: '#main',
    data: {
      userid: "",
      auth_code: "",
      team: "",
    },
    created: function(){
      var cookie = this.getCookie();
      this.userid = cookie["userid"];
      this.auth_code = cookie["code"];
      console.log(this.userid);
      console.log(this.auth_code);
    },
    watch: {
    },
    methods: {
        sendJoinTeam : function(team) {
            var sendInfo = {};
            sendInfo.userid = this.userid;
            sendInfo.team = team;

            fetch("/team/" + this.userid + "?code=" + this.auth_code, {
                method: 'post',
                mode: 'cors',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(sendInfo)
            }).then(response => {
                if (response.ok) {
                    location.href='/client/' + this.userid
                } else {
                    console.log("NG");
                    throw new Error();
                }
            });
        },
        getCookie : function() {
            var arr = new Array();
            if(document.cookie != ''){
                var tmp = document.cookie.split('; ');
                for(var i=0; i<tmp.length; i++){
                    var data = tmp[i].split('=');
                    arr[data[0]] = decodeURIComponent(data[1]);
                }
            }
            return arr;
        },
    }
  })
}
