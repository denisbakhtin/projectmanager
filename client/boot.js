import './css/site.scss'
import m from 'mithril'
import 'moment'
import users from './components/users/users'
import login from './components/account/login'
import logout from './components/account/logout'
import roles from './components/roles/roles'
import register from './components/account/register'
import activationNotice from './components/account/activation_notice'
import activate from './components/account/activate'
import reset from './components/account/reset'
import resetNotice from './components/account/reset_notice'
import manage from './components/account/manage'
import home from './components/home'

import projects from './components/projects/projects'
import new_project from './components/projects/new_project'
import edit_project from './components/projects/edit_project'
import project from './components/projects/project'

import statuses from './components/statuses/statuses'
import new_status from './components/statuses/new_status'
import edit_status from './components/statuses/edit_status'
import status from './components/statuses/status'

import tasksteps from './components/tasksteps/tasksteps'
import new_taskstep from './components/tasksteps/new_taskstep'
import edit_taskstep from './components/tasksteps/edit_taskstep'
import taskstep from './components/tasksteps/taskstep'

import tasks from './components/tasks/tasks'
import new_task from './components/tasks/new_task'
import edit_task from './components/tasks/edit_task'
import task from './components/tasks/task'

import layout from './components/shared/layout'
import public_layout from './components/shared/layout_public'
import Auth from './utils/auth'

var app_root = document.getElementById("app-root");

var routes = {
    '/': route(home, false),
    '/login': route(login, false),
    '/register': route(register, false),
    '/activation_notice': route(activationNotice, false),
    '/activate/:token': route(activate, false),
    '/reset': route(reset, false),
    '/reset/:token': route(reset, false),
    '/reset_notice': route(resetNotice, false),
    '/manage': route(manage),
    '/roles': route(roles),
    '/users': route(users),

    '/projects': route(projects),
    '/projects/new': route(new_project),
    '/projects/edit/:id': route(edit_project),
    '/projects/:id': route(project),

    '/statuses': route(statuses),
    '/statuses/new': route(new_status),
    '/statuses/edit/:id': route(edit_status),
    '/statuses/:id': route(status),

    '/task_steps': route(tasksteps),
    '/task_steps/new': route(new_taskstep),
    '/task_steps/edit/:id': route(edit_taskstep),
    '/task_steps/:id': route(taskstep),

    '/tasks': route(tasks),
    '/tasks/new': route(new_task),
    '/tasks/edit/:id': route(edit_task),
    '/tasks/:id': route(task),

    '/logout': route(logout, false),
};

//route wrapper 
function route(comp, requiresAuth = true) {
    if (requiresAuth)
        return {
            onmatch: checkAuthorized(comp)
        }
    else
        return withLayout(comp)
}

//render component with layout
function withLayout(comp) {
    return {
        view: () => Auth.isLoggedIn() ? m(layout, comp) : m(public_layout, comp)
    }
}

//Authorization route filter
function checkAuthorized(comp) {
    return function (args, path) {
        if (!Auth.isLoggedIn()) {
            localStorage.returnURL = path
            m.route.set("/login")
        } else return withLayout(comp)
    }
}

m.route(app_root, "/", routes)

//turn on webpack hot reload
if (module.hot) {
    module.hot.accept()
}