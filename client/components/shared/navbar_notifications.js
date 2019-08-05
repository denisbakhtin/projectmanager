import m from 'mithril'

const Notifications = {
    view(vnode) {
        return m('li.nav-item.dropdown.mr-2#navbar-notifications', [
            m('a.nav-link.dropdown-toggle[href=#]', [
                m('span.fa.fa-bell-o'),
                m('span.badge.badge-pill.badge-primary', 5)
            ])
        ])
    }
}
export default Notifications;