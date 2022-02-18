import { Component } from 'react'

import Slider from './components/Slider';
import Preview from './components/Preview';

import './App.css';

class App extends Component {
  constructor(props) {
    super(props)
    this.state = {
      src: '',

      previewSrc: null,
      loading: false,
      height: 0,
      width: 0,

      scale: 0,
      x: 0,
      y: 0,
    };

    this.img = null;
  }

  updateOverlay(src) {
    this.setState({src})
    try {
      new URL(src);
    } catch(e) {
      if(e instanceof TypeError) {
        this.img = null;
        this.setState({previewSrc: null, loading: false});
        return;
      }
      throw e;
    }

    this.setState({ loading: true });

    this.img = document.createElement('img');
    this.img.onload = () => {
      this.setState({
        loading: false,
        previewSrc: this.img.src,
        height: this.img.height,
        width: this.img.width,
      });
    };
    this.img.onerror = () => {
      this.setState({loading: false, previewSrc: null});
    }
    this.img.src = src;
  }

  renderSlider(name) {
    return (
      <Slider
        name={name}
        value={this.state[name]}
        onUpdate={(value) => {
          const newState = {};
          newState[name] = parseInt(value, 10);
          this.setState(newState);
        }}
      />
    )
  }

  render() {
    return (
      <div className="App">
        <div className="controls">
          <label>Overlay URL</label>
          <input
            type="text"
            id="src"
            name="src"
            value={this.state.src}
            onChange={(e) => this.updateOverlay(e.target.value)}
          />

          <label>Scale</label>
          {this.renderSlider('scale')}

          <label>Horizontal Offset</label>
          {this.renderSlider('x')}
          
          <label>Vertical Offset</label>
          {this.renderSlider('y')}
        </div>

        <Preview
          src={this.state.previewSrc}
          loading={this.state.loading}
          scale={this.state.scale}
          x={this.state.x}
          y={this.state.y}
          width={this.state.width}
          height={this.state.height}
        />
      </div>
    )
  }
}

export default App
