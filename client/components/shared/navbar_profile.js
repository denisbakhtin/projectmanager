import m from 'mithril'
import Auth from '../../utils/auth'
import dropdown from './dropdown'

export default function Profile() {
    return {
        view: (vnode) => {
            return m(dropdown, {
                children: [
                    m('a.nav-link#navbar-profile[href=#]', [
                        m('img.img-round.mr-1[width=23px][height=23px]', {
                            src: '/public/images/user_icon.png'
                        }),
                        m('span', Auth.getAuthenticatedUser().name)
                    ]),
                    m('.dropdown-menu', [
                        m('a.dropdown-item[href=#!/manage]', 'Account'),
                        m('a.dropdown-item.text-danger[href=#!/logout]', 'Logout'),
                    ]),
                ]
            })
        }
    }
}
