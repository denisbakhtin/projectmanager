import m from 'mithril'
import Auth from '../utils/auth'
import login from './account/login';
import dashboard from './dashboard/dashboard'

export default function Home() {
    return {
        view(vnode) {
            return Auth.isLoggedIn() ? m(dashboard) : m(login);
        }
    }
}
