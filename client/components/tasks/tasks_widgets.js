import m from 'mithril'
import service from '../../utils/service.js'
import { responseErrors, humanSpent } from '../../utils/helpers'
import { startTask } from '../shared/active_task'
import error from '../shared/error'
//import moment from 'moment'

export function TasksCountWidget() {
    let count = 0,
        errors,

        get = () =>
            service.getTasksSummary()
                .then((result) => count = result.count)
                .catch((error) => errors = responseErrors(error))

    return {
        oninit(vnode) {
            get()
        },

        view(vnode) {

            return m(".card.count-widget",
                m('a.card-body[href=#!/tasks]', [
                    m('.count', count),
                    m('.description', 'Tasks'),
                    (errors) ? m('i.fa.fa-exclamation-circle.error-icon', { title: responseErrors(errors) }) : null,
                ])
            )
        }
    }
}

export function LatestTasksWidget() {
    let latestTasks = [],
        errors = null,

        get = () =>
            service.getLatestTasks()
                .then((result) => latestTasks = result)
                .catch((error) => errors = responseErrors(error))

    return {
        oninit(vnode) {
            get()
        },

        view(vnode) {
            return m(".card.task-widget",
                m('.card-body', [
                    m('.widget-title', [
                        "Recently created tasks",
                        m('a.ml-auto[href=#!/tasks]', 'Show all')
                    ]),
                    m('table.table', [
                        m('thead', [
                            m('tr', [
                                m('th', 'Name'),
                                m('th.shrink.text-center', 'Start')
                            ])
                        ]),
                        m('tbody', [
                            (latestTasks && latestTasks.length > 0) ?
                                latestTasks.map((task) => m('tr', [
                                    m('td', m('a', { href: '#!/tasks/' + task.id }, task.name)),
                                    m('td.buttons.shrink.text-center',
                                        m('button.btn.btn-icon.btn-primary', { onclick: () => startTask(task, () => get()) }, m('i.fa.fa-play'))
                                    ),
                                ])) : m('tr', m('td.text-center[colspan=2]', 'The list is empty'))
                        ])
                    ]),
                    m(error, { errors: responseErrors(errors) }),
                ])
            )
        }
    }
}

export function LatestTaskLogsWidget() {
    let latestTaskLogs = [],
        errors = null,

        get = () =>
            service.getLatestTaskLogs()
                .then((result) => latestTaskLogs = result)
                .catch((error) => errors = responseErrors(error))

    return {
        oninit(vnode) {
            get()
        },

        view(vnode) {
            return m(".card.task-widget",
                m('.card-body', [
                    m('.widget-title', [
                        "Recently run tasks",
                        m('a.ml-auto[href=#!/reports/spent]', 'Show all')
                    ]),
                    m('table.table', [
                        m('thead', [
                            m('tr', [
                                m('th', 'Name'),
                                //m('th.shrink.text-center', 'When'),
                                m('th.shrink.text-center', 'Duration')
                            ])
                        ]),
                        m('tbody', [
                            (latestTaskLogs && latestTaskLogs.length > 0) ?
                                latestTaskLogs.map((log) => m('tr', [
                                    m('td', m('a', { href: '#!/tasks/' + log.task.id }, log.task.name)),
                                    //m('td.shrink.text-center', moment(log.created_at).fromNow()),
                                    m('td.shrink.text-center', humanSpent(log.minutes)),
                                ])) : m('tr', m('td.text-center[colspan=3]', 'The list is empty'))
                        ])
                    ]),
                    m(error, { errors: responseErrors(errors) }),
                ])
            )
        }
    }
}