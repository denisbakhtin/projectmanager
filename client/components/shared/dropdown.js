import m from 'mithril'

function Dropdown() {
    let show = false
    let key = 'dropdown'
    let loadState = () => {
        let s = sessionStorage.getItem(key)
        show = (s) ? JSON.parse(s) : show
    }
    let storeState = (val) => {
        show = val
        sessionStorage.setItem(key, JSON.stringify(val))
    }
    return {
        oninit: (vnode) => {
            key = vnode.attrs.id || key
            loadState()
        },
        view: (vnode) => {
            for (let child of vnode.attrs.children) {
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
                onclick: (e) => storeState(!show),
            }, vnode.attrs.children)
        }
    }
}

export default Dropdown