import m from 'mithril'
import dropdown from './dropdown'

const SideMenu = {
    view(vnode) {
        return m('#side-menu.wide', [
            m('ul.nav.flex-column', [
                m('li.nav-item', m('a.nav-link[href=#!/]', [
                    m('span.fa.fa-home'),
                    m('span', 'Dashboard')
                ])),
                m('li.nav-item', m('a.nav-link[href=#!/projects]', [
                    m('span.fa.fa-list'),
                    m('span', 'Projects')
                ])),
                m('li.nav-item', m('a.nav-link[href=#!/tasks]', [
                    m('span.fa.fa-check-square-o'),
                    m('span', 'Tasks')
                ])),
                m('li.nav-item', m('a.nav-link[href=#!/users]', [
                    m('span.fa.fa-users'),
                    m('span', 'Users')
                ])),
                m('li.nav-item', m('a.nav-link[href=#!/manage]', [
                    m('span.fa.fa-key'),
                    m('span', 'Manage Account')
                ])),
                m(dropdown, {
                    id: 'settings-dropdown',
                    children: [
                        m('a.nav-link[href=#]', [
                            m('span.fa.fa-wrench'),
                            m('span', 'Settings')
                        ]),
                        m('.dropdown-menu', [
                            m('a.dropdown-item[href=#!/statuses]', 'Project Statuses'),
                            m('a.dropdown-item[href=#!/task_steps]', 'Task Steps'),
                            m('a.dropdown-item[href=#!/roles]', 'User Roles'),
                        ]),
                    ]
                }),
                m('li.nav-item', m('a.nav-link[href=#!/logout]', [
                    m('span.fa.fa-sign-out'),
                    m('span', 'Logout')
                ])),
            ]),
            m('.bottom-links.row', [
                m('a.d-inline-block.col[href=#]', m('span.fa.fa-lock')),
                m('a.d-inline-block.col[href=#]', m('span.fa.fa-download')),
                m('a.d-inline-block.col[href=#]', m('span.fa.fa-globe')),
                m('a.d-inline-block.col[href=#]', m('span.fa.fa-phone')),
            ])
        ])
    }
}

export default SideMenu;