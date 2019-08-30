import m from 'mithril'
import error from '../shared/error'
import state from './state'

const Statuses = {
    oninit(vnode) {
        state.errors = []
        state.getAll()
    },
    
    view(vnode) {
        let ui = vnode.state
        return m(".statuses", [
            m('h1.mb-4', 'Project status'),
            m('table.table', [
                m('thead', [
                    m('tr', [
                        m('th[scope=col]', 'Name'),
                        m('th[scope=col]', 'Description'),
                        m('th[scope=col]', 'Order'),
                        m('th.shrink.text-center[scope=col]', 'Actions')
                    ])
                ]),
                m('tbody', [
                    state.statuses ? 
                        state.statuses.map((status) => {
                            return m('tr', { key: status.id }, [
                                m('td', status.name),
                                m('td', status.description),
                                m('td', status.order),
                                m('td.shrink.text-center', m('button.btn.btn-outline-primary.btn-sm[type=button]', { onclick: () => { m.route.set('/statuses/edit/'+status.id) } }, m('i.fa.fa-pencil')))
                            ])
                        }) : null
                ])
            ]),
            state.errors.length ? m(error, { errors: state.errors }) : null,
            m('.actions.mt-4', [
                m('button.btn.btn-primary[type=button]', { onclick: () => { m.route.set('/statuses/new') } }, "New status")
            ]),
        ])
    }
}

export default Statuses;