import m from 'mithril'
import Auth from '../../utils/auth'

const Logout = {
    oninit(vnode) {
        Auth.logout()
    },
    view(vnode) {
        return null
    }
}

export default Logout;