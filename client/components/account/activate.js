import m from 'mithril'
import Auth from '../../utils/auth'
import error from '../shared/error'

const Activate = {
    view(vnode) {
        return m('.activation-notice', [
            m('h1', "Account activation"),
            m('p', "Click on the button below to finish registration."),
            m(error, { errors: Auth.errors }),
            m('button.btn.btn-primary[type=button]', { onclick: () => {Auth.activate(m.route.param('token'))} }, "Activate"),
        ])
    }
}

export default Activate