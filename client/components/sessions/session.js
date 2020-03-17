import m from 'mithril'
import {
    groupLogsByProject,
    responseErrors,
    humanTaskSpent,
    humanProjectSpent,
    humanAllProjectsSpent,
} from '../../utils/helpers'
import error from '../shared/error'
import service from '../../utils/service.js'

export default function Session() {
    let session,
        projects = [],
        logs = [],
        id,
        errors = [],

        get = (id) =>
            service.getSession(id)
                .then((result) => {
                    session = result
                    logs = session.task_logs.slice(0)
                    projects = groupLogsByProject(logs)
                })
                .catch((error) => errors = responseErrors(error))

    return {
        oninit(vnode) {
            id = m.route.param('id')
            get(id)
        },

        view(vnode) {
            return m(".session", [
                m('h1.title.mb-4', 'Session #' + id),
                (session && session.contents && session.contents.length > 0) ?
                    m('p.session-contents', session.contents) : null,
                projects.length > 0 ? [
                    projects.map((project) =>
                        m('.session-box.dashboard-box.mb-4', { key: project.id }, [
                            m('h5.strong', m('a', { href: '#!/projects/' + project.id }, project.name)),
                            m('table.table', [
                                m('thead', m('tr', [
                                    m('th', 'Task'),
                                    m('th.shrink', 'Spent')
                                ])),
                                m('tbody', project.tasks.map((task) =>
                                    m('tr', { key: task.id }, [
                                        m('td', task.name),
                                        m('td', humanTaskSpent(task))
                                    ])
                                )),
                                m('tfoot', m('tr', [
                                    m('th', 'Total spent'),
                                    m('th', humanProjectSpent(project))
                                ])),
                            ]),
                        ])
                    ),

                    m('.dashboard-box.spent-total',
                        m('table.w-100',
                            m('tr', [
                                m('td', 'Total spent on all projects'),
                                m('td.shrink', m('mark.ml-2', humanAllProjectsSpent(projects, true))),
                            ])
                        )
                    ),
                ] : m('p.alert.alert-info', 'This session is empty.'),
                m(error, { errors: errors }),
                m('.actions.mt-4', [
                    m('button.btn.btn-outline-secondary.mr-2[type=button]', {
                        onclick: () => window.history.back()
                    }, "Back"),
                ])
            ])
        }
    }
}