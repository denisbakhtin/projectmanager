import m from 'mithril'
import service from '../../utils/service.js'
import {
    addSuccess
} from '../shared/notifications'
import {
    responseErrors,
    groupLogsByProject,
    humanTaskSpent,
    humanProjectSpent,
    humanSessionSpent,
} from '../../utils/helpers'
import error from '../shared/error'

export default function Session() {
    let session = { task_logs: [] },
        commitedTasks = {},
        errors = [],
        projects = [],
        logs = [],

        setContents = (contents) => session.contents = contents,

        taskCommited = (task) => commitedTasks[task.id] != undefined,
        commitTask = (task) => {
            commitedTasks[task.id] = task
            task.task_logs.forEach((log) => session.task_logs.push(log))
        },
        unCommitTask = (task) => {
            delete commitedTasks[task.id]
            task.task_logs.forEach((log) => {
                let ind = session.task_logs.indexOf(log)
                if (ind > -1) session.task_logs.splice(ind, 1)
            })
        },

        projectCommited = (project) => project.tasks.every((task) => commitedTasks[task.id] != undefined),
        commitProject = (project) => {
            project.tasks.forEach((task) => {
                commitedTasks[task.id] = task
                task.task_logs.forEach((log) => session.task_logs.push(log))
            })
        },
        unCommitProject = (project) => {
            project.tasks.forEach((task) => {
                delete commitedTasks[task.id]
                task.task_logs.forEach((log) => {
                    let ind = session.task_logs.indexOf(log)
                    if (ind > -1) session.task_logs.splice(ind, 1)
                })
            })
        },

        //requests
        newSession = () =>
            service.newSession()
                .then((result) => {
                    logs = result.slice(0)
                    projects = groupLogsByProject(logs)
                }).catch((error) => errors = responseErrors(error)),


        create = () => {
            //clear circular references
            session.task_logs.forEach((log) => delete log.task)
            service.createSession(session)
                .then((result) => {
                    addSuccess("Session created.")
                    m.route.set('/sessions')
                })
                .catch((error) => errors = responseErrors(error))
        }

    return {
        oninit(vnode) {
            newSession()
        },
        view(vnode) {
            return m(".session", (session !== { task_logs: [] }) ? [
                m('h1.title', 'New session'),
                m('.form-group', [
                    m('label', "Session comment"),
                    m('textarea.form-control', {
                        oninput: (e) => setContents(e.target.value),
                        value: session.contents
                    })
                ]),

                projects.length > 0 ? [
                    projects.map((project) =>
                        m('.session-box.dashboard-box.mb-4', { key: project.id }, [
                            m('h5.strong', m('a', { href: '#!/projects/' + project.id }, project.name)),
                            m('table.table', [
                                m('thead', m('tr', [
                                    m('th', 'Task'),
                                    m('th.shrink', 'Spent'),
                                    m('th.shrink', m('.custom-control.custom-switch', [
                                        m('input.custom-control-input[type=checkbox]', {
                                            id: 'project-' + project.id,
                                            checked: projectCommited(project),
                                            onchange: (val) => projectCommited(project) ? unCommitProject(project) : commitProject(project),
                                        }),
                                        m('label.custom-control-label', { for: 'project-' + project.id })
                                    ]))
                                ])),
                                m('tbody', project.tasks.map((task) =>
                                    m('tr', { key: task.id }, [
                                        m('td', m('a', { href: '#!/tasks/' + task.id }, task.name)),
                                        m('td.shrink', humanTaskSpent(task)),
                                        m('td.shrink', m('.custom-control.custom-switch', [
                                            m('input.custom-control-input[type=checkbox]', {
                                                id: 'task-' + task.id,
                                                checked: taskCommited(task),
                                                onchange: (val) => taskCommited(task) ? unCommitTask(task) : commitTask(task),
                                            }),
                                            m('label.custom-control-label', { for: 'task-' + task.id })
                                        ]))
                                    ])
                                )),
                                m('tfoot', m('tr', [
                                    m('th', 'Total spent'),
                                    m('th.shrink', humanProjectSpent(project)),
                                    m('th.shrink')
                                ])),
                            ]),
                        ])
                    ),

                    m('.dashboard-box.spent-total',
                        m('table.w-100',
                            m('tr', [
                                m('td', 'Total spent on all projects'),
                                m('td.shrink', humanSessionSpent(session, true) ? m('mark.ml-2', humanSessionSpent(session, true)) : ''),
                            ])
                        )
                    ),
                ] : m('p.alert.alert-info', 'No uncommited task activities found.'),

                m('.my-2', m(error, { errors: errors })),

                m('.actions.mt-4', [
                    m('button.btn.btn-primary.mr-2[type=button]', {
                        onclick: create
                    }, [
                        m("i.fa.fa-check.mr-1"),
                        "Submit"
                    ]),
                    m('button.btn.btn-outline-secondary[type=button]', {
                        onclick: () => window.history.back()
                    }, "Cancel")
                ]),
            ] : m('Loading...'))
        }
    }
}
