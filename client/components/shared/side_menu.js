import m from 'mithril'
import dropdown from './dropdown'
import state from './state'

const SideMenu = {
    view(vnode) {
        return m('#side-menu', {class: state.sidebarCollapsed ? 'narrow' : 'wide'}, [
            m('ul.nav.flex-column', [
                m('li.nav-item', m('a.nav-link[href=#!/]', [
                    m('span.fa.fa-home'),
                    m('span.title', 'Dashboard')
                ])),
                m('li.nav-item', m('a.nav-link[href=#!/projects]', [
                    m('span.fa.fa-list'),
                    m('span.title', 'Projects')
                ])),
                m('li.nav-item', m('a.nav-link[href=#!/tasks]', [
                    m('span.fa.fa-check-square-o'),
                    m('span.title', 'Tasks')
                ])),
                m('li.nav-item', m('a.nav-link[href=#!/users]', [
                    m('span.fa.fa-users'),
                    m('span.title', 'Users')
                ])),
                m('li.nav-item', m('a.nav-link[href=#!/manage]', [
                    m('span.fa.fa-key'),
                    m('span.title', 'Manage Account')
                ])),
                m(dropdown, {
                    id: 'settings-dropdown',
                    keepState: true,
                    children: [
                        m('a.nav-link[href=#]', [
                            m('span.fa.fa-wrench'),
                            m('span.title', 'Settings')
                        ]),
                        m('.dropdown-menu', [
                            m('a.dropdown-item[href=#!/statuses]', 'Project Statuses'),
                            m('a.dropdown-item[href=#!/task_steps]', 'Task Steps'),
                            m('a.dropdown-item[href=#!/roles]', 'User Roles'),
                            m('a.dropdown-item[href=#!/settings]', 'Site Settings'),
                        ]),
                    ]
                }),
                m('li.nav-item', m('a.nav-link[href=#!/support]', [
                    m('span.fa.fa-question-circle'),
                    m('span.title', 'Contact Support')
                ])),
                m('li.nav-item', m('a.nav-link[href=#!/logout]', [
                    m('span.fa.fa-sign-out'),
                    m('span.title', 'Logout')
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
