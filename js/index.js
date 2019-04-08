window.onload = function () {
  var app = new Vue({
    el: '#main',
    data: {
      is_join: {},
      is_survival: {},
      count: 0,
    },
    created: function(){
        setInterval(() => { this.count++ }, 1000)
        this.getSurvival();
        this.getJoin();
    },
    watch: {
        count: function(){
        },
    },
    methods: {
        getSurvival : function() {
            fetch("/survival/me", {
                method: 'GET',
                mode: 'cors',
            }).then((response) => {
                return response.json();
            }).then((json) => {
                this.is_survival = json;
            });
        },
        getJoin : function() {
            fetch("/join/me", {
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
