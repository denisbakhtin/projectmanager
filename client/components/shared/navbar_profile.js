import m from 'mithril'
import Auth from '../../utils/auth'

const Profile = {
    view(vnode) {
        return m('li.nav-item.dropdown.mr-2#navbar-profile', [
            m('a.nav-link.dropdown-toggle[href=#]', [
                m('img.img-round.mr-1[width=23px][height=23px]', {
                    src: '/public/images/user_icon.png'
                }),
                m('span', Auth.getAuthenticatedUser().name)
            ])
        ])
    }
}
export default Profile;