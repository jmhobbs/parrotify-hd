import './Toggle.css';

function Toggle(props) {
  return (
    <div className="toggle-wrapper">
      <label htmlFor={props.name}>{props.label}</label>
      <label className="toggle" htmlFor={props.name}>
        <input
          type="checkbox"
          name={props.name}
          id={props.name}
          checked={props.checked ? 'checked' : ''}
          onChange={(e) => props.onUpdate(e.target.checked)}
        />
        <div className="slider"></div>
      </label>
    </div>
  )
}

export default Toggle;
