import m from 'mithril'
import error from '../shared/error'
import service from '../../utils/service.js'
import projects_item from '../projects/projects_item.js'
import tasks_item from '../tasks/tasks_item.js'
import comments_item from '../comments/comments_item.js'
import categories_item from '../categories/categories_item.js'
import global_state from '../shared/state'
import {
    responseErrors
} from '../../utils/helpers'

export default function Search() {
    let projects = [],
        tasks = [],
        comments = [],
        categories = [],
        errors = [],

        search = () =>
            (global_state.query) ?
                service.getSearch(global_state.query)
                    .then((result) => {
                        categories = result.categories.slice(0)
                        projects = result.projects.slice(0)
                        tasks = result.tasks.slice(0)
                        comments = result.comments.slice(0)
                    })
                    .catch((error) => errors = responseErrors(error))
                : null

    return {
        oninit: (vnode) => search(),

        view(vnode) {
            return m(".projects.tasks.comments", [
                (categories && categories.length > 0) ? [
                    m('h3.mb-3', 'Categories'),
                    m('ul.dashboard-box.box-list.mb-3', [
                        categories.map((cat) => m(categories_item, { key: cat.id, category: cat, onUpdate: search }))
                    ]),
                ] : null,
                (projects && projects.length > 0) ? [
                    m('h3.mb-3', 'Projects'),
                    m('ul.dashboard-box.box-list.mb-3',
                        projects.map((proj) => m(projects_item, { key: proj.id, project: proj, onUpdate: search }))
                    )
                ] : null,
                (tasks && tasks.length > 0) ? [
                    m('h3.mb-3', 'Tasks'),
                    m('ul.dashboard-box.box-list.mb-3',
                        tasks.map((task) => m(tasks_item, { key: task.id, task: task, onUpdate: search }))
                    )
                ] : null,
                (comments && comments.length > 0) ? [
                    m('h3.mb-3', 'Comments'),
                    m('ul.dashboard-box.box-list.mb-3',
                        comments.map((comment) => m(comments_item, { key: comment.id, comment: comment, onUpdate: search }))
                    ),
                ] : null,
                (projects.length == 0 && tasks.length == 0 && comments.length == 0 && categories.length == 0) ? m('p', 'Search query returned no results.') : null,
                m(error, { errors: errors }),
                m('.actions.mt-4', [
                    m('button.btn.btn-outline-secondary[type=button]', {
                        onclick: () => window.history.back()
                    }, [
                        "Back"
                    ])
                ]),
            ])
        }
    }
}
