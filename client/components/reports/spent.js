import m from 'mithril'
import {
    groupLogsByProject,
    responseErrors,
    humanTaskSpent,
    humanProjectSpent,
    humanAllProjectsSpent,
    firstOfWeek,
    minusDays,
    firstOfMonth
} from '../../utils/helpers'
import error from '../shared/error'
import service from '../../utils/service.js'

export default function SpentReport() {
    let projects = [],
        logs = [],
        errors = [],
        activeFilter = 'All',

        get = () =>
            service.getSpent()
                .then((result) => logs = result.slice(0))
                .catch((error) => errors = responseErrors(error)),

        filterByDates = (dateFrom, dateTo) => {
            let minDate = new Date(1980, 1, 1),
                maxDate = new Date(2100, 1, 1)

            if (logs) {
                let tlogs = logs.filter((log) => Date.parse(log.created_at) >= (dateFrom ?? minDate) && Date.parse(log.created_at) < (dateTo ?? maxDate))
                projects = groupLogsByProject(tlogs)
            }
        }

    return {
        oninit(vnode) {
            get().then((res) => filterByDates(null, null))
        },

        view(vnode) {
            return m(".spent-report", [
                m('h1.title.mb-4', 'Time Spent Report'),
                m('.filters', [
                    m('button.btn.btn-link', {
                        class: (activeFilter == 'All') ? 'active' : '',
                        onclick: () => { activeFilter = 'All'; filterByDates(null, null) }
                    }, 'All'),
                    m('button.btn.btn-link', {
                        class: (activeFilter == 'Week') ? 'active' : '',
                        onclick: () => { activeFilter = 'Week'; filterByDates(firstOfWeek(), Date.now()) }
                    }, 'This week'),
                    m('button.btn.btn-link', {
                        class: (activeFilter == '7days') ? 'active' : '',
                        onclick: () => { activeFilter = '7days'; filterByDates(minusDays(7), Date.now()) }
                    }, '7 days'),
                    m('button.btn.btn-link', {
                        class: (activeFilter == 'Month') ? 'active' : '',
                        onclick: () => { activeFilter = 'Month'; filterByDates(firstOfMonth(), Date.now()) }
                    }, 'This month'),
                    m('button.btn.btn-link', {
                        class: (activeFilter == '30days') ? 'active' : '',
                        onclick: () => { activeFilter = '30days'; filterByDates(minusDays(30), Date.now()) }
                    }, '30 days'),
                ]),
                projects.length > 0 ? [
                    projects.map((project) =>
                        m('.spent-report-box.dashboard-box.mb-4', { key: project.id }, [
                            m('h5.strong', m('a', { href: '#!/projects/' + project.id }, project.name)),
                            m('table.table', [
                                m('thead', m('tr', [
                                    m('th', 'Task'),
                                    m('th.shrink', 'Spent')
                                ])),
                                m('tbody', project.tasks.map((task) =>
                                    m('tr', { key: task.id }, [
                                        m('td', m('a', { href: '#!/tasks/' + task.id }, task.name)),
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
                ] : m('p.alert.alert-info', 'No task activities for the specified period found.'),
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