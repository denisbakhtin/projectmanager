import m from 'mithril'

/* function Dropdown() {
    let show = false
    console.log('hehhhhh')
    return {
        view(vnode) {
            if (vnode.children.length > 0)
                for (let child of vnode.children) {
                    //prevent default navigation to /#
                    if (child.tag == "a")
                        child.attrs['onclick'] = (e) => e.preventDefault()
                    //prevent event propagation to parent item which leads to dropdown collapse
                    if (child.tag == "div" && child.children.length > 0)
                        for (let item of child.children) {
                            item.attrs['onclick'] = (e) => e.stopPropagation()
                        }
                }

            return m('li.nav-item.dropdown', {
                class: show ? 'show' : '',
                onclick: (e) => show = !show,
            }, vnode.children)
        }
    }
}
 */
const Dropdown = {
    show: false,
    oninit: (vnode) => console.log('oninit', vnode.state.show),
    oncreate: (vnode) => console.log('oncreate', vnode.state.show),
    view(vnode) {
        if (vnode.children.length > 0)
            for (let child of vnode.children) {
                //prevent default navigation to /#
                if (child.tag == "a")
                    child.attrs['onclick'] = (e) => e.preventDefault()
                //prevent event propagation to parent item which leads to dropdown collapse
                if (child.tag == "div" && child.children.length > 0)
                    for (let item of child.children) {
                        item.attrs['onclick'] = (e) => e.stopPropagation()
                    }
            }

        return m('li.nav-item.dropdown', {
            class: vnode.state.show ? 'show' : '',
            onclick: (e) => vnode.state.show = !vnode.state.show,
        }, vnode.children)
    }
}

export default Dropdown