import './Preview.css';

import parrot from '../assets/parrot.png';

const DEFAULT_SCALE = 70;
const STILL_IMAGE_CENTER = [42, 62];

function Preview(props) {
  if(props.loading) {
    return (<div className='preview'>Loading...</div>)
  }

  if(props.src === null) {
    return (<div className='preview' />)
  }

  const width = DEFAULT_SCALE + props.scale;
  const height = Math.floor((props.height / props.width) * width);
  const left = Math.floor(STILL_IMAGE_CENTER[0] - (width/2) + props.x);
  const top = Math.floor(STILL_IMAGE_CENTER[1] - (height/2) + props.y);

  const download_url = '/parrot.gif?' + [
    `src=${encodeURIComponent(props.src)}`,
    `scale=${props.scale}`,
    `x=${props.x}`,
    `y=${props.y}`,
    `flip=${props.flip}`,
    `rotate=${props.rotate}`
  ].join('&');

  const style = {
    width: width + 'px',
    height: height + 'px',
    top: top + 'px',
    left: left + 'px',
  };

  style['-webkit-transform'] = 'rotate(' + props.rotate + 'deg)';
  style.transform = 'rotate(' + props.rotate + 'deg)';

  if(props.flip) {
    style['-webkit-transform'] += ' scaleX(-1)';
    style.transform += ' scaleX(-1)';
  }

  return (
    <>
    <div className='preview'>
      <img className="parrot" src={parrot} />
      <img
        className="overlay"
        src={props.src}
        style={style}
      />
    </div>
        <a className='button' href={download_url} target='_blank'>Download GIF</a>
    </>
  )
}

export default Preview;
