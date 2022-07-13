
export default function ClientCreate() {
  return (
    <div>
      <form>
        <input id="name" placeholder="Name" maxLength="100" />
        <textarea id="description" placeholder="Description" maxLength="300" />
        <button>Create</button>
      </form>
    </div>
  )
}