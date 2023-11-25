import { Meta, StoryObj } from '@storybook/react'
import { Footer } from './Footer'

//👇 This default export determines where your story goes in the story list
const meta: Meta<typeof Footer> = {
  component: Footer,
}

export default meta
type Story = StoryObj<typeof Footer>

export const FirstStory: Story = {
  args: {
    primary: false,
    backgroundColor: 'bg-purple-500',
    label: 'First',
  },
}

export const SecondStory: Story = {
  args: {
    primary: false,
    size: 'large',
    backgroundColor: 'bg-amber-600',
    label: 'Second',
  },
}
