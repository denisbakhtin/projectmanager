import m from 'mithril'

const Support = {
    view(vnode) {
        return m('li.nav-item.dropdown.mr-2#navbar-support', [
            m('a.nav-link[href=#]', [
                m('span.fa.fa-comments-o')
            ])
        ])
    }
}
export default Support