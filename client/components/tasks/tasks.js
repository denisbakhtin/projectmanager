import m from 'mithril'
import error from '../shared/error'
import {
    responseErrors
} from '../../utils/helpers'
import service from '../../utils/service'


export default function Tasks() {
    let tasks = [],
        errors = [],

        getAll = () =>
        service.getTasks()
        .then((result) => tasks = result.slice(0))
        .catch((error) => errors = responseErrors(error))

    return {
        oninit(vnode) {
            getAll()
        },

        view(vnode) {
            return m(".tasks", [
                m('h1.mb-4', 'Project tasks'),
                m('table.table', [
                    m('thead', [
                        m('tr', [
                            m('th[scope=col]', 'Name'),
                            m('th[scope=col]', 'State'),
                            m('th[scope=col]', 'Description'),
                            m('th[scope=col]', 'Assigned User'),
                            m('th.shrink.text-center[scope=col]', 'Actions')
                        ])
                    ]),
                    m('tbody', [
                        tasks ?
                        tasks.map((task) => {
                            return m('tr', {
                                key: task.id
                            }, [
                                m('td', task.name),
                                m('td', task.task_step.name),
                                m('td', task.description),
                                m('td', task.project_user.user.name),
                                m('td.shrink.text-center', m('button.btn.btn-outline-primary.btn-sm[type=button]', {
                                    onclick: () => {
                                        m.route.set('/tasks/edit/' + task.id)
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
                            m.route.set('/tasks/new')
                        }
                    }, "New task")
                ]),
            ])
        }
    }
}
