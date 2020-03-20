import m from 'mithril'
import service from '../../utils/service.js'
import { responseErrors, humanSpent } from '../../utils/helpers'
import { startTask } from '../shared/active_task'
//import moment from 'moment'

const state = {
    count: 0,
    latestTasks: [],
    latestTaskLogs: [],
    errors: null,
    lastRun: 0,

    //methods
    get: () => {
        //10 sec threshold
        if (Math.floor((Date.now() - state.lastRun) / 1000) > 10) {
            state.lastRun = Date.now()
            service.getTasksSummary()
                .then((result) => {
                    state.errors = null
                    state.count = result.count
                    state.latestTasks = result.latest_tasks
                    state.latestTaskLogs = result.latest_task_logs
                })
                .catch((error) => state.errors = responseErrors(error))
        }
    }

}

export function TasksCountWidget() {
    return {
        oninit(vnode) {
            state.get()
        },

        view(vnode) {

            return m(".card.count-widget",
                m('a.card-body[href=#!/tasks]', [
                    m('.count', state.count),
                    m('.description', 'Tasks'),
                    (state.errors) ? m('i.fa.fa-exclamation-circle.error-icon', { title: responseErrors(state.errors) }) : null,
                ])
            )
        }
    }
}

export function LatestTasksWidget() {
    return {
        oninit(vnode) {
            state.get()
        },

        view(vnode) {
            return m(".card.task-widget",
                m('.card-body', [
                    m('.widget-title', "Recently created tasks"),
                    m('table.table', [
                        m('thead', [
                            m('tr', [
                                m('th', 'Name'),
                                m('th.shrink.text-center', 'Start')
                            ])
                        ]),
                        m('tbody', [
                            (state.latestTasks && state.latestTasks.length > 0) ?
                                state.latestTasks.map((task) => m('tr', [
                                    m('td', m('a', { href: '#!/tasks/' + task.id }, task.name)),
                                    m('td.buttons.shrink.text-center',
                                        m('button.btn.btn-icon.btn-primary', { onclick: () => startTask(task, () => state.get()) }, m('i.fa.fa-play'))
                                    ),
                                ])) : m('tr', m('td[colspan=2]', 'The list is empty'))
                        ])
                    ]),
                    (state.errors) ? m('.error', responseErrors(state.errors)) : null,
                ])
            )
        }
    }
}

export function LatestTaskLogsWidget() {
    return {
        oninit(vnode) {
            state.get()
        },

        view(vnode) {
            return m(".card.task-widget",
                m('.card-body', [
                    m('.widget-title', "Recently run tasks"),
                    m('table.table', [
                        m('thead', [
                            m('tr', [
                                m('th', 'Name'),
                                //m('th.shrink.text-center', 'When'),
                                m('th.shrink.text-center', 'Duration')
                            ])
                        ]),
                        m('tbody', [
                            (state.latestTaskLogs && state.latestTaskLogs.length > 0) ?
                                state.latestTaskLogs.map((log) => m('tr', [
                                    m('td', m('a', { href: '#!/tasks/' + log.task.id }, log.task.name)),
                                    //m('td.shrink.text-center', moment(log.created_at).fromNow()),
                                    m('td.shrink.text-center', humanSpent(log.minutes)),
                                ])) : m('tr', m('td[colspan=3]', 'The list is empty'))
                        ])
                    ]),
                    (state.errors) ? m('.error', responseErrors(state.errors)) : null,
                ])
            )
        }
    }
}