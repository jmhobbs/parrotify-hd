import Preview from '../components/Preview';

import zol from './assets/zol.png';

export default {
  title: 'Component/Preview',
  component: Preview,
  args: {
    src: null,
    loading: false,
    width: 0,
    height: 0,
    scale: 0,
    x: 0,
    y: 0,
  }
};

const Template = (args) => <Preview {...args} />;

export const Default = Template.bind({});

export const Loading = Template.bind({});
Loading.args = {
  loading: true,
};

export const Loaded = Template.bind({});
Loaded.args = {
  src: zol,
  width: 324,
  height: 442,
};

export const Scaled = Template.bind({});
Scaled.args = {
  src: zol,
  width: 324,
  height: 442,
  scale: -20,
};

export const Shifted = Template.bind({});
Shifted.args = {
  src: zol,
  width: 324,
  height: 442,
  x: -20,
  y: 20,
};