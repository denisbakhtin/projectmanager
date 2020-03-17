import m from 'mithril'
import error from '../shared/error'
import service from '../../utils/service.js'
import {
    responseErrors,
    ISODateToHtml5
} from '../../utils/helpers'
import projects_item from '../projects/projects_item'
import tasks_item from '../tasks/tasks_item'

export default function Category() {
    let category = {},
        id,
        errors = [],

        get = () =>
            service.getCategory(id)
                .then((result) => category = result)
                .catch((error) => errors = responseErrors(error))

    return {
        oninit(vnode) {
            id = m.route.param('id')
            get()
        },

        view(vnode) {
            return m(".category", (category) ? [
                m('h1.title.mb-2', [
                    category.name,
                    m('button.btn.btn-default.ml-2[type=button]', {
                        onclick: () => m.route.set('/categories/edit/' + category.id)
                    }, "Edit"),
                ]),
                (category.projects && category.projects.length > 0) ? [
                    m('h4', 'Projects'),
                    m('ul.dashboard-box.box-list.mb-4',
                        category.projects.map((proj) => m(projects_item, { key: proj.id, project: proj, onUpdate: get }))
                    ),
                ] : null,
                (category.tasks && category.tasks.length > 0) ? [
                    m('h4', 'Tasks'),
                    m('ul.dashboard-box.box-list',
                        category.tasks.map((task) => m(tasks_item, { key: task.id, task: task, onUpdate: get }))
                    )
                ] : null,
                m(error, { errors: errors }),
                m('.actions.mt-4', [
                    m('button.btn.btn-outline-secondary.mr-2[type=button]', {
                        onclick: () => window.history.back()
                    }, "Back"),
                ])
            ] : m('Loading...'))
        }
    }
}
