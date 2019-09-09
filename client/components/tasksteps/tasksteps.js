import m from 'mithril'
import error from '../shared/error'
import {
    responseErrors
} from '../../utils/helpers'

export default function TaskSteps() {
    let steps = [],
        errors = [],

        getAll = () =>
        service.getTaskSteps()
        .then((result) => steps = result.slice(0))
        .catch((error) => errors = responseErrors(error))

    return {
        oninit(vnode) {
            getAll()
        },

        view(vnode) {
            return m(".task_steps", [
                m('h1.mb-4', 'Task steps'),
                m('table.table', [
                    m('thead', [
                        m('tr', [
                            m('th[scope=col]', 'Name'),
                            m('th[scope=col]', 'Is final'),
                            m('th[scope=col]', 'Order'),
                            m('th.shrink.text-center[scope=col]', 'Actions')
                        ])
                    ]),
                    m('tbody', [
                        steps ?
                        steps.map((step) => {
                            return m('tr', {
                                key: step.id
                            }, [
                                m('td', step.name),
                                m('td', step.is_final),
                                m('td', step.ord),
                                m('td.shrink.text-center', m('button.btn.btn-outline-primary.btn-sm[type=button]', {
                                    onclick: () => {
                                        m.route.set('/task_steps/edit/' + step.id)
                                    }
                                }, m('i.fa.fa-pencil')))
                            ])
                        }) : null
                    ])
                ]),
                errors.length ? m(error, {
                    errors: errors
                }) : null,
                m('.actions.mt-4', [
                    m('button.btn.btn-primary[type=button]', {
                        onclick: () => {
                            m.route.set('/task_steps/new')
                        }
                    }, "New step")
                ]),
            ])
        }
    }
}
