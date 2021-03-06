﻿import m from 'mithril'
import Auth from '../../utils/auth'
import error from '../shared/error'

export default function Login() {
    return {
        oninit(vnode) {
            Auth.errors = []
        },
        view(vnode) {
            let ui = vnode.state;
            return m("#auth-form-wrapper", [
                m('.card', [
                    m('.card-body', [
                        m('h3.card-title.text-center', "User login"),
                        m('form[action=#]', [
                            m('.form-group', [
                                m('label.control-label', "Email"),
                                m('input.form-control[type=email][name=email][placeholder="Enter your email"]', {
                                    oninput: (e) => Auth.setEmail(e.target.value),
                                    value: Auth.email
                                })
                            ]),
                            m('.form-group', [
                                m('label.control-label', "Password"),
                                m('input.form-control[type=password][name=email][placeholder="Enter your password"]', {
                                    oninput: (e) => Auth.setPassword(e.target.value),
                                    value: Auth.password
                                })
                            ]),
                            m('button.btn.btn-primary.btn-lg.mb-2[type=submit]', {
                                onclick: Auth.login
                            }, "Login"),
                            m(error, { errors: Auth.errors })
                        ]),
                        Auth.errors.length > 0 ? m('.text-center', m('a[href=#!/reset]', "Forgot your password?")) : null,
                        m('.text-center', m('a[href=#!/register]', "Don't have an account?"))
                    ])
                ])
            ])
        }
    }
}
