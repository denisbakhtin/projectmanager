import m from 'mithril'
import jwt_decode from 'jwt-decode'
import {
    emailIsValid,
    responseErrors
} from './helpers'

const Auth = {
    name: '',
    email: '',
    currentPassword: '',
    password: '',
    passwordConfirm: '',
    setName: (name) => Auth.name = name,
    setEmail: (email) => Auth.email = email.toLowerCase(),
    setCurrentPassword: (currentPassword) => Auth.currentPassword = currentPassword,
    setPassword: (password) => Auth.password = password,
    setPasswordConfirm: (passwordConfirm) => Auth.passwordConfirm = passwordConfirm,
    errors: [],
    tokenNotExpired: () => !!localStorage.expires && (localStorage.expires > Date.now() / 1000),
    isLoggedIn: () => !!localStorage.token && Auth.tokenNotExpired(),
    isAdmin: () => !!localStorage.token && (localStorage.role === "Admin"),
    login: () => {
        if (Auth.validateLogin())
            m.request({
                method: "POST",
                url: "/api/login",
                body: {
                    email: Auth.email,
                    password: Auth.password
                }
            }).then((result) => {
                Auth.storeTokenData(result.token)
                let returnURL = localStorage.returnURL || '/'
                delete localStorage.returnURL
                m.route.set(returnURL) //redirect to home page or returnURL
            }).catch((error) => Auth.errors = responseErrors(error))
        return false
    },
    register: () => {
        if (Auth.validateRegister())
            m.request({
                method: "POST",
                url: "/api/register",
                body: {
                    name: Auth.name,
                    email: Auth.email,
                    password: Auth.password
                }
            }).then((result) => {
                //m.route.set('/activation_notice')
                Auth.storeTokenData(result.token)
                m.route.set('/')
            }).catch((error) => Auth.errors = responseErrors(error))
        return false
    },
    /*
    activate(activationToken) {
        m.request({
            method: "POST",
            url: "/api/activate",
            body: { token: activationToken }
        }).then((result) => {
            Auth.storeTokenData(result.token)
            m.route.set('/') //redirect to home page
        }).catch((error) => Auth.errors = responseErrors(error))
    },
    */
    update: () => {
        if (Auth.validateUpdate())
            m.request({
                method: "PUT",
                url: "/api/account",
                body: {
                    name: Auth.name,
                    email: Auth.email,
                    current_password: Auth.currentPassword,
                    new_password: Auth.password
                },
                headers: {
                    Authorization: Auth.authHeader()
                }
            }).then((result) => {
                Auth.storeTokenData(result.token)
                m.route.set('/') //redirect to home page
            }).catch((error) => Auth.errors = responseErrors(error))
    },
    requestReset: () => {
        if (Auth.validateResetRequest())
            m.request({
                method: "POST",
                url: "/api/forgot",
                body: {
                    email: Auth.email
                }
            }).then((result) => {
                m.route.set('/reset_notice')
            }).catch((error) => Auth.errors = responseErrors(error))
    },
    reset: (token) => {
        if (Auth.validateReset())
            m.request({
                method: "POST",
                url: "/api/reset",
                body: {
                    token: token,
                    password: Auth.password
                }
            }).then((result) => {
                Auth.storeTokenData(result.token)
                m.route.set('/') //redirect to home page
            }).catch((error) => Auth.errors = responseErrors(error))
    },
    logout: () => {
        delete localStorage.token
        delete localStorage.user_id
        delete localStorage.name
        delete localStorage.email
        delete localStorage.role
        delete localStorage.expires
        m.route.set('/') //redirect to home page
    },
    authHeader: () => (!!localStorage.token) ? 'Bearer ' + localStorage.token : null,
    storeTokenData: (token) => {
        let decoded = jwt_decode(token)
        localStorage.token = token
        localStorage.user_id = decoded.user_id
        localStorage.name = decoded.name
        localStorage.email = decoded.sub
        localStorage.role = decoded.role
        localStorage.expires = decoded.exp
    },
    getAuthenticatedUser: () =>
        (!!localStorage.token) ? {
            user_id: localStorage.user_id,
            name: localStorage.name,
            email: localStorage.email,
            token: localStorage.token,
            role: localStorage.role,
            expires: localStorage.expires
        } : null,
    validateLogin: () => {
        Auth.errors = []
        if (!Auth.email || !emailIsValid(Auth.email))
            Auth.errors.push("Valid email is required.")
        if (!Auth.password)
            Auth.errors.push("Password is required.")
        return Auth.errors.length == 0
    },
    validateRegister: () => {
        Auth.errors = []
        if (!Auth.name) Auth.errors.push("Name is required.")
        if (!Auth.email || !emailIsValid(Auth.email))
            Auth.errors.push("Valid email is required.")
        if (!Auth.password || Auth.password.length < 8)
            Auth.errors.push("Password must be atleast 8 characters.")
        if (Auth.password !== Auth.passwordConfirm)
            Auth.errors.push("Password does not match the confirm password.")
        return Auth.errors.length == 0
    },
    validateUpdate: () => {
        Auth.errors = []
        if (!Auth.name) Auth.errors.push("Name is required.")
        if (!Auth.email || !emailIsValid(Auth.email))
            Auth.errors.push("Valid email is required.")
        if (Auth.currentPassword) {
            if (!Auth.password || Auth.password.length < 8)
                Auth.errors.push("Password must be atleast 8 characters.")
            if (Auth.password !== Auth.passwordConfirm)
                Auth.errors.push("Password does not match the confirm password.")
        }
        return Auth.errors.length == 0
    },
    validateResetRequest: () => {
        Auth.errors = []
        if (!Auth.email || !emailIsValid(Auth.email))
            Auth.errors.push("Valid email is required.")
        return Auth.errors.length == 0
    },
    validateReset: () => {
        Auth.errors = []
        if (!Auth.password || Auth.password.length < 8)
            Auth.errors.push("Password must be atleast 8 characters.")
        if (Auth.password !== Auth.passwordConfirm)
            Auth.errors.push("Password does not match the confirm password.")
        return Auth.errors.length == 0
    }
}

export default Auth
