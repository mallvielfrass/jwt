const app = new Vue({

    delimiters: ['${', '}'],
    el: '#app',
    data: {
        errors: [],
        login: null,
        email: null,
        movie: null,
        passF: null,
        passLogin: null,
        passT: null,
        logreg: false,
        RegSuccess:true,
        SuccFalse:true,
        SuccFalseMessage:""
    },
    methods: {
        axiosLog() {
            url= `/api/auth/create_session?login=${app.login}&password=${app.passF}`
            console.log("(axiosReg) send: ",url)
            axios.get( url)
                .then(function (response) {
                    // handle success
                    console.log(response);
                    if (response.data.Error!=""){
                        console.log(response.data.Error);

                        document.cookie = `login=${response.data.Login}; expires=Thu, 01 Jan 1970 00:00:00 UTC;`;
                        document.cookie = `session=${response.data.Session}; expires=Thu, 01 Jan 1970 00:00:00 UTC;`;
                    }else{
                        console.log(response.data.Login);
                        console.log(response.data.Session);
                        console.log(response.data.Expiry);
                        dtime=new Date(response.data.Expiry)
                        console.log(dtime)
                        document.cookie = `login=${response.data.Login}; expires=${ dtime};`;
                        document.cookie = `session=${response.data.Session}; expires=${ dtime};`;
                    }
                   // document.cookie = `login=${app.email}; expires=30d`;
                    //document.cookie = `password=${app.passF}; expires=30d`;
                    // $cookies.set("login", app.email,{ expires: "30d" } );
                    // $cookies.set("password", app.passF,{ expires: "30d" } );
                    
                })
                .catch(function (error) {
                    // handle error
                    console.log(error);
                });
        },
        axiosReg() {
            url= `/api/auth/create_user?login=${app.login}&password=${app.passF}`
            console.log("(axiosReg) send: ",url)
            axios.get( url)
                .then(function (response) {
                    // handle success
                    console.log(response);
                    if (response.data.Error!=""){
                        console.log(response.data.Error);

                        document.cookie = `login=${response.data.Login}; expires=Thu, 01 Jan 1970 00:00:00 UTC;`;
                        document.cookie = `session=${response.data.Session}; expires=Thu, 01 Jan 1970 00:00:00 UTC;`;
                    }else{
                        console.log(response.data.Login);
                        console.log(response.data.Session);
                        console.log(response.data.Expiry);
                        dtime=new Date(response.data.Expiry)
                        console.log(dtime)
                        document.cookie = `login=${response.data.Login}; expires=${ dtime};`;
                        document.cookie = `session=${response.data.Session}; expires=${ dtime};`;
                    }
                   // document.cookie = `login=${app.email}; expires=30d`;
                    //document.cookie = `password=${app.passF}; expires=30d`;
                    // $cookies.set("login", app.email,{ expires: "30d" } );
                    // $cookies.set("password", app.passF,{ expires: "30d" } );
                    
                })
                .catch(function (error) {
                    // handle error
                    console.log(error);
                });

        },
        logHandler() {
            this.errors = [];
            console.log("this.login: " + this.login + " this.passF: " + this.passF )
            if (!this.login) {
                this.errors.push('Укажите имя.');
            }
           
            if (this.passF.length == 0) {
                this.errors.push('укажите пароль');
            }

            if (this.errors.length == 0) {
                console.log("log ok")
                this.axiosLog()
            }

        },
        subHandler() {
            this.errors = [];
            console.log("this.login: " + this.login + " this.passF: " + this.passF + " this.passT: " + this.passT)
            if (!this.login) {
                this.errors.push('Укажите имя.');
            }
            if (this.passF != this.passT) {
                this.errors.push('пароли не совпадают');
            }
            if (this.passF.length == 0) {
                this.errors.push('укажите пароль');
            }
            if (this.errors.length == 0) {
                console.log("reg ok")
                this.RegSuccess=false
                this.axiosReg()
            }

        },
    }

})