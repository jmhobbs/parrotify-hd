import Slider from '../components/Slider';

export default {
  title: 'Component/Slider',
  component: Slider,
  args: {
    value: 0,
  }
};

const Template = (args) => <Slider {...args} />;

export const Default = Template.bind({});

export const Preset = Template.bind({});
Preset.args = {
  value: 42,
};
