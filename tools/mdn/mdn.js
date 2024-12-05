// This is a user script which generates constants from the MDN Web Docs.
// Headers: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers
// Status:  https://developer.mozilla.org/en-US/docs/Web/HTTP/Status

(() => {
// noo you can't just change the prototype of objects in JS!!!
HTMLElement.prototype.qsa = function(query) { return Array.from(this.querySelectorAll(query)) }
HTMLElement.prototype.qs = function(query) { return this.querySelector(query) }

const capitalize = str => str?.charAt(0)?.toUpperCase?.() + str?.slice?.(1)
const pair_reducer = (acc, el, i, arr) => i % 2 === 0 ? [...acc, [el, arr[i + 1]]] : acc

let reporting_endpoints_counter = 0 // can't be bothered to implement finer logic.

const child_query = "h2+div dl :is(dt, dd)"
const group_query = "section:has(h2+div dl :is(dt a, dd))"
const supported_types = ["Header", "Status"]

const source_url = document.body.qs("#on-github a").href
const contributors = document.body.qs(".last-modified-date a").href
const timestamp = document.body.qs(".last-modified-date time").innerText

const generated_type = (v => v == "Headers" ? "Header" : v)(window.location.href.split("/").pop())

if (!supported_types.includes(generated_type)) throw new Error(`unsupported type: ${generated_type}`)

class Entry {
	constructor(dt, dd) {
		this.dtel = dt
		this.ddel = dd
	}

	get ident() {
		let split = this.raw.split(" ")
		let val = split.length >= 2 ? split.splice(1).join("_").replaceAll("'", "") : this.raw
		let res = val.replaceAll("-", "_")
		if (res == "Reporting_Endpoints") {
			reporting_endpoints_counter += 1
			if (reporting_endpoints_counter == 2) {
				res = "SSE_" + res
			}
		}
		
		return capitalize(res)
	}

	get link() {
		let href = this.dtel.qs("a").getAttribute("href")
		href ??= `/en-US/docs/Web/HTTP/Headers#${this.raw}`
		return `https://developer.mozilla.org${href}`
	}

	get comment() {
		let lines = this.ddel.innerText.split(/\s+/).reduce((acc, word) => {
			let lastLine = acc[acc.length - 1]
			if ((lastLine + word).length > 80) {
				acc.push(word)
				return acc
			}
			acc[acc.length - 1] = lastLine + (lastLine ? " " : "") + word
			
			return acc
		}, [`\t// ${this.ident}: `])
		.join(`\n\t// `);

		return `${lines}\n\t// ${this.link}`
	}

	get value() {
		let split = this.raw.split(" ")
		let val = split.length >= 2 ? split[0] : `"${this.raw}"`
		return val
	}

	get raw() {
		return this.dtel.qs("a code").innerText
	}

	get format() {
		return `${this.comment}\n\t${this.ident} ${generated_type} = ${this.value}`
	}
}

class Group {
	constructor(el) { this.el = el }

	get children() {
		return this.el
			.qsa(child_query)
			.reduce(pair_reducer, [])
			.map(a => new Entry(a[0], a[1]));
	}

	get comment() {
		let {innerText, href} = this.el.qs("a")
		return `// ${innerText}\n// ${href}`
	}

	get format() {
		return `
${this.comment}
const (
${this.children.map(a => a.format).join("\n")}
)
`
	}
}


const repo_url = "https://github.com/mdn/content"

const header = `// This file is generated directly from the from the MDN Web Docs.
// The script formats the licensed prose content to code comments.
// See the script responsible for this in the docs/js directory of this repository.
// The documentation in this file is available under the CC-BY-SA 2.5 license,
// as is all prose content on MDN Web Docs.
//
// Attribution:
// - Source: MDN Web Docs (https://developer.mozilla.org/)
// - License: CC-BY-SA 2.5 (https://creativecommons.org/licenses/by-sa/2.5/)
// - Contributors: ${contributors}
// - Source file:  ${source_url}
// - Source repo:  ${repo_url}
//
// At time of generation, the source file was last modified on ${timestamp}.
package hnet`


let res = document.body.qsa(group_query)
	.map(a => new Group(a))
	.reduce((acc, cur) => `${acc}${cur.format}`, "")
	.replaceAll("  ", " ");


console.log(`${header}\n${res}`)
})();
