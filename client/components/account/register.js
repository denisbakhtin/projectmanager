import m from 'mithril'
import Auth from '../../utils/auth'
import error from '../shared/error'

const Register = {
    oninit(vnode) {
        Auth.errors = []
        Auth.name = "Denis Bakhtin"
        Auth.email = "denis.bakhtin@gmail.com"
        Auth.password = "12345678"
        Auth.passwordConfirm = "12345678"
    },
    view(vnode) {
        let ui = vnode.state;
        return m("#auth-form-wrapper", [
            m('.card', [
                m('.card-body', [
                    m('h3.card-title.text-center', "User registration"),
                    m('form', [
                        m('.form-group', [
                            m('label.control-label', "Name"),
                            m('input.form-control[placeholder="Enter your name"]', {
                                oninput: function (e) {
                                    Auth.setName(e.target.value)
                                },
                                value: Auth.name
                            })
                        ]),
                        m('.form-group', [
                            m('label.control-label', "Email"),
                            m('input.form-control[type=email][placeholder="Enter your email"]', {
                                oninput: function (e) {
                                    Auth.setEmail(e.target.value)
                                },
                                value: Auth.email
                            })
                        ]),
                        m('.form-group', [
                            m('label.control-label', "Password"),
                            m('input.form-control[type=password][placeholder="Enter your password"]', {
                                oninput: function (e) {
                                    Auth.setPassword(e.target.value)
                                },
                                value: Auth.password
                            })
                        ]),
                        m('.form-group', [
                            m('label.control-label', "Confirm password"),
                            m('input.form-control[type=password][placeholder="Confirm your password"]', {
                                oninput: function (e) {
                                    Auth.setPasswordConfirm(e.target.value)
                                },
                                value: Auth.passwordConfirm
                            })
                        ]),
                        m('button.btn.btn-primary.btn-lg.mb-2[type=submit]', {
                            onclick: Auth.register
                        }, "Register"),
                        m(error, {
                            errors: Auth.errors
                        })
                    ]),
                    m('.text-center', m('a[href=#!/login]', "Already have an account?"))
                ])
            ])
        ])
    }
}

export default Register;