function Slider(props) {
  return (
    <div>
      <button onClick={() => props.onUpdate(props.value - 1)}>-</button>
      <input 
        type="range"
        min="-64"
        max="64"
        value={props.value}
        id={props.name}
        name={props.name}
        onChange={(e) => props.onUpdate(e.target.value)}
        />
      <button onClick={() => props.onUpdate(props.value + 1)}>+</button>
    </div>
  )
}

export default Slider;
