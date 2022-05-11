import Toggle from '../components/Toggle';

export default {
  title: 'Component/Toggle',
  component: Toggle,
  args: {
    name: 'flip',
    label: 'Flip Image',
    checked: false,
  }
};

const Template = (args) => <Toggle {...args} />;

export const Default = Template.bind({});

export const On = Template.bind({});
On.args = {
  checked: true,
};
