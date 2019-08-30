import m from 'mithril'
import Auth from '../../utils/auth'
import error from '../shared/error'

const Manage = {
    oninit(vnode) {
        let user = Auth.getAuthenticatedUser()
        Auth.name = user.name
        Auth.email = user.email
        Auth.currentPassword = ''
        Auth.password = ''
        Auth.passwordConfirm = ''
        Auth.errors = []
    },
    view(vnode) {
        let ui = vnode.state;
        return m("#auth-form-wrapper", [
            m('.card', [
                m('.card-body', [
                    m('h3.card-title.text-center', "Account management"),
                    m('form', [
                        m('.form-group', [
                            m('label.control-label', "Name"),
                            m('input.form-control[placeholder="Enter your name"]', {
                                oninput: (e) => Auth.setName(e.target.value),
                                value: Auth.name
                            })
                        ]),
                        m('.form-group', [
                            m('label.control-label', "Email"),
                            m('input.form-control[type=email][placeholder="Enter your email"]', {
                                oninput: (e) => Auth.setEmail(e.target.value),
                                value: Auth.email
                            })
                        ]),
                        m('.form-group', [
                            m('label.control-label', "Current password"),
                            m('input.form-control[type=password][placeholder="Enter your current password"]', {
                                oninput: (e) => Auth.setCurrentPassword(e.target.value),
                                value: Auth.currentPassword
                            })
                        ]),
                        m('.form-group', [
                            m('label.control-label', "Password"),
                            m('input.form-control[type=password][placeholder="Enter your password"]', {
                                oninput: (e) => Auth.setPassword(e.target.value),
                                value: Auth.password
                            })
                        ]),
                        m('.form-group', [
                            m('label.control-label', "Confirm password"),
                            m('input.form-control[type=password][placeholder="Confirm your password"]', {
                                oninput: (e) => Auth.setPasswordConfirm(e.target.value),
                                value: Auth.passwordConfirm
                            })
                        ]),
                        m('button.btn.btn-primary.btn-lg.mb-2[type=submit]', {
                            onclick: Auth.update
                        }, "Update"),
                        m(error, {
                            errors: Auth.errors
                        })
                    ])
                ])
            ])
        ])
    }
}

export default Manage;