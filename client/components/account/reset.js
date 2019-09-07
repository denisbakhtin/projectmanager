import m from 'mithril'
import Auth from '../../utils/auth'
import error from '../shared/error'

export default function Reset() {
  return {
    oninit(vnode) {
        Auth.errors = []
        Auth.password = ''
        Auth.passwordConfirm = ''
    },
    view(vnode) {
        let token = m.route.param('token')
        return m('.password-reset', [
            m('h1', "Password reset"),
            token ? [
                m('.form-group', [
                    m('label.control-label', "New password"),
                    m('input.form-control[type=password][placeholder="Enter your new password"]', {
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
                m(error, {
                    errors: Auth.errors
                }),
                m('button.btn.btn-primary[type=button]', {
                    onclick: () => {
                        Auth.reset(m.route.param('token'))
                    }
                }, "Activate"),
            ] : [
                m('.form-group', [
                    m('label', "Please, enter your email"),
                    m('input.form-control[type=email]', {
                        oninput: (e) => Auth.setEmail(e.target.value),
                        value: Auth.email
                    })
                ]),
                m(error, {
                    errors: Auth.errors
                }),
                m('button.btn.btn-primary[type=button]', {
                    onclick: Auth.requestReset
                }, "Continue"),
            ]
        ])
    }
  }
}
