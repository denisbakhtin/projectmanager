import m from 'mithril'
import error from '../shared/error'
import service from '../../utils/service.js'
import projects_item from './projects_item.js'

const Filters = Object.freeze({
    ALL: (project) => true,
    OPEN: (project) => project.archived == false,
    ARCHIVED: (project) => project.archived == true,
    FAVORITE: (project) => project.favorite == true,
})

export default function Projects() {
    let projects = [],
        errors = [],
        filter = Filters.ALL,

        activeClass = (fil) => (filter === fil) ? "active" : "",

        getAll = () =>
            service.getProjects()
                .then((result) => projects = result.slice(0))
                .catch((error) => errors = responseErrors(error))

    return {
        oninit(vnode) {
            getAll()
        },

        view(vnode) {
            let filteredProjects = projects.filter(filter)

            return m(".projects", [
                m('h1.title', 'Projects'),
                m('.filters', [
                    m('button.btn.btn-link', { class: activeClass(Filters.ALL), onclick: () => filter = Filters.ALL }, "All"),
                    m('button.btn.btn-link', { class: activeClass(Filters.OPEN), onclick: () => filter = Filters.OPEN }, "Open"),
                    m('button.btn.btn-link', { class: activeClass(Filters.ARCHIVED), onclick: () => filter = Filters.ARCHIVED }, "Archived"),
                    m('button.btn.btn-link', { class: activeClass(Filters.FAVORITE), onclick: () => filter = Filters.FAVORITE }, "Favorite"),
                ]),
                (filteredProjects && filteredProjects.length > 0) ? m('ul.dashboard-box.box-list',
                    filteredProjects.map((proj) => m(projects_item, { key: proj.id, project: proj, onUpdate: getAll }))
                ) : m('p.empty-list', 'The list is empty'),
                m(error, { errors: errors }),
                m('.actions.mt-4', [
                    m('button.btn.btn-primary[type=button]', {
                        onclick: () => m.route.set('/projects/new')
                    }, [
                        m('i.fa.fa-plus.mr-1'),
                        "New project"
                    ])
                ]),
            ])
        }
    }
}
